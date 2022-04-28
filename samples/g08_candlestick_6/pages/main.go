package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"syscall/js"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

		x := float64(i+1) * (bar.w * 1.2)
		ctx.Call("fillRect", x-0.5, high, 1, low)
		ctx.Call("fillRect", x-bar.w/2.0, top, bar.w, bottom)
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

func drawPriceLabel(ctx js.Value, size Size, jsonData CandleStickResponse) {
	ctx.Set("fillStyle", "gray")
	ctx.Set("font", "12px Roboto")

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

	tick := 10.
	priceRange := max - min

	fmt.Println(max, min, priceRange)

	tickRange := float64(priceRange) / tick
	yRatio := float64(max-min) / (size.h * 0.9)

	// adjust
	cnt := 0
	p := tickRange
	for p >= 10. {
		p = p / 10.
		cnt++
	}
	base, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", p), 64)
	tickRange = base * float64(math.Pow10(cnt))

	if tickRange < 1 {
		tickRange = 1
	}
	lowerBound := tickRange * float64(int(1+min/int(tickRange)))
	upperBound := tickRange * float64(int(1+max/int(tickRange)))
	printer := message.NewPrinter(language.English)
	for i := lowerBound; i <= upperBound; i += tickRange {
		ctx.Call("save")
		ctx.Call("translate", 0, (float64(max)-i)/yRatio)
		ctx.Call("fillText", printer.Sprintf("%.0f", i), 30, 0)
		ctx.Call("restore")
	}
}

func main() {
	size := Size{w: 800, h: 500}
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
	h2.Set("innerHTML", "Chart06 WebAssembly! (made with GO)")
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

	frameSize := Size{w: size.w - 100.0, h: size.h}
	ctx.Call("save")
	ctx.Call("translate", 100, 0)
	ctx.Call("strokeRect", 0, 0, frameSize.w, frameSize.h*0.9)
	drawCandle(ctx, frameSize, jsonData)
	drawDateLabel(ctx, frameSize, jsonData)
	ctx.Call("restore")

	drawPriceLabel(ctx, size, jsonData)

	body.Set("onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			event := args[0]
			if !event.Get("target").Equal(canvas) {
				return nil
			}
			x := event.Get("clientX").Float()
			y := event.Get("clientY").Float()

			js.Global().Get("window").Call("alert", fmt.Sprintf("x=%f, y=%f", x, y))
			return nil
		}))

	<-make(chan struct{})
}
