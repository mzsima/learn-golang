package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"syscall/js"
	"time"
)

type CandleStickResponse struct {
	Status       int
	Responsetime string
	Data         []Candlestick
}

type Candlestick struct {
	OpenTime int     `json:",string"`
	Open     int     `json:",string"`
	High     int     `json:",string"`
	Low      int     `json:",string"`
	Close    int     `json:",string"`
	Volume   float64 `json:",string"`
}

type Size struct {
	w float64
	h float64
}

func drawCandle(ctx js.Value, size Size, jsonData CandleStickResponse) {
	max := 0
	min := int(^uint(0) >> 1)
	for _, v := range jsonData.Data {
		if max < v.High {
			max = v.High
		}
		if min > v.Low {
			min = v.Low
		}
	}

	yRatio := float64(max-min) / (size.h * 0.9)

	dx := size.w / float64(len(jsonData.Data))
	bar := Size{w: dx * 0.8, h: 0}
	for i, v := range jsonData.Data {

		fmt.Println(v.Open, v.High, v.Low, v.Close)

		high := float64(max-v.High) / yRatio
		low := float64(v.High-v.Low) / yRatio

		var top, bottom float64
		if v.Open > v.Close {
			ctx.Set("fillStyle", "#e32739")
			top = float64(max-v.Open) / yRatio
			bottom = float64(v.Open-v.Close) / yRatio
		} else {
			ctx.Set("fillStyle", "#5eb142")
			top = float64(max-v.Close) / yRatio
			bottom = float64(v.Close-v.Open) / yRatio
		}

		fmt.Println(top, high, low, bottom)
		ctx.Call("fillRect", float64(i+1)*(bar.w*1.2)-0.5, high, 1, low)
		ctx.Call("fillRect", float64(i+1)*(bar.w*1.2)-bar.w/2.0, top, bar.w, bottom)
	}
}

func drawDateLabel(ctx js.Value, size Size, jsonData CandleStickResponse) {
	dx := size.w / float64(len(jsonData.Data))

	ctx.Set("fillStyle", "gray")

	for i, v := range jsonData.Data {
		if i%2 != 0 {
			continue
		}

		hhmm := time.UnixMilli(int64(v.OpenTime)).Format("15:04")
		ctx.Call("save")
		ctx.Call("translate", float64(i)*dx+8, size.h*0.92)
		ctx.Call("rotate", math.Pi*0.4)
		ctx.Call("fillText", hhmm, 0, 0)
		ctx.Call("restore")
	}
}

func main() {
	size := Size{w: 500, h: 500}
	scale := js.Global().Get("devicePixelRatio").Float()
	doc := js.Global().Get("document")
	if !doc.Truthy() {
		panic("unable to get 'document object")
	}

	body := doc.Get("body")
	if !body.Truthy() {
		panic("unable to get 'body' object")
	}

	h2 := doc.Call("createElement", "h2")
	h2.Set("innerHTML", "Chart02 WebAssembly! (made with GO)")
	body.Call("appendChild", h2)

	canvas := doc.Call("createElement", "canvas")
	ctx := canvas.Call("getContext", "2d")
	canvas.Set("width", size.w*scale)
	canvas.Set("height", size.h*scale)
	body.Call("appendChild", canvas)

	canvas.Get("style").Set("width", size.w)
	canvas.Get("style").Set("height", size.h)

	ctx.Call("scale", scale, scale)
	ctx.Set("fillStyle", "rgba(128, 128, 128, 0.5)")
	ctx.Call("beginPath")

	resp, err := http.Get("/kline")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var jsonData CandleStickResponse
	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		log.Fatalln(err)
	}
	ctx.Call("strokeRect", 0, 0, size.w, size.h*0.9)

	drawCandle(ctx, size, jsonData)
	drawDateLabel(ctx, size, jsonData)
}
