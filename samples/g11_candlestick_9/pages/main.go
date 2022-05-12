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
	Path     js.Value
}

type Size struct {
	w float64
	h float64
}

func (c CandleStickResponse) Candles(size Size) (candles []Candlestick) {
	max := 0
	min := int(^uint(0) >> 1)

	for _, v := range c.Data {
		if max < v.High {
			max = v.High
		}
		if min > v.Low {
			min = v.Low
		}
	}

	yRatio := float64(max-min) / (size.h * 0.9)
	dx := size.w / float64(len(c.Data))
	bar := Size{w: dx * 0.8, h: 0}

	for i, v := range c.Data {
		high := float64(max-v.High) / yRatio
		low := float64(v.High-v.Low) / yRatio

		var top, bottom float64
		if v.Open > v.Close {
			top = float64(max-v.Open) / yRatio
			bottom = float64(v.Open-v.Close) / yRatio
		} else {
			top = float64(max-v.Close) / yRatio
			bottom = float64(v.Close-v.Open) / yRatio
		}
		x := float64(i+1)*(bar.w*1.2) + 100
		path := js.Global().Get("Path2D").New()
		path.Call("rect", x-0.5, high, 1, low)
		path.Call("rect", x-bar.w/2.0, top, bar.w, bottom)
		v.Path = path
		candles = append(candles, v)
	}

	return
}

func (c Candlestick) Draw(ctx js.Value) {
	if c.Open > c.Close {
		ctx.Set("fillStyle", "#e32739")
	} else {
		ctx.Set("fillStyle", "#5eb142")
	}
	ctx.Call("fill", c.Path)
}

func (c Candlestick) Hit(x, y, ctx js.Value) {
	// ? isPointInPathは canvasのscaleに影響をうけてる？
	isHit := ctx.Call("isPointInPath", c.Path, x.Float()*2, y.Float()*2)
	if isHit.Bool() {
		js.Global().Call("alert", fmt.Sprintf("High: %d, Low: %d", c.High, c.Low))
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
		ctx.Call("translate", float64(i)*dx+108, size.h*0.92)
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
	h2.Set("innerHTML", "Chart09 WebAssembly! (made with GO)")
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
	ctx.Call("strokeRect", 100, 0, frameSize.w, frameSize.h*0.9)
	candles := jsonData.Candles(frameSize)
	for _, v := range candles {
		v.Draw(ctx)
	}
	drawDateLabel(ctx, frameSize, jsonData)
	drawPriceLabel(ctx, size, jsonData)

	canvas.Set("onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			event := args[0]
			if !event.Get("target").Equal(canvas) {
				return nil
			}
			x := event.Get("offsetX")
			y := event.Get("offsetY")
			for _, v := range candles {
				v.Hit(x, y, ctx)
			}
			return nil
		}))

	<-make(chan struct{})
}
