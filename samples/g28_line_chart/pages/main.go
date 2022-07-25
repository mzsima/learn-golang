package main

import (
	"fmt"
	"math"
	"syscall/js"
)

type Position struct {
	x float64
	y float64
}

type Size struct {
	w float64
	h float64
}

type MyCanvas struct {
	canvas js.Value
	ctx    js.Value
	center Position
}

type MyData struct {
	label  string
	values []int
}

func (m MyCanvas) drawLineChart(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Call("translate", 0, 160)
	ctx.Call("scale", 1, -1) // origin to lower left corner
	for _, d := range data {
		for i, v := range d.values {
			if i == 0 {
				ctx.Call("beginPath")
			}
			x := i*40 + 40
			ctx.Call("lineTo", x, v)
			ctx.Call("stroke")
		}

		for i, v := range d.values {
			ctx.Call("beginPath")
			x := i*40 + 40
			if d.label == "A" {
				ctx.Call("arc", x, v, 3, 0, math.Pi*2.0)
			} else {
				ctx.Call("rect", x-3, v-3, 6, 6)
			}
			ctx.Call("fill")
		}
	}
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Set("font", "12px Roboto")
	for i := 0; i <= 100; i += 20 {
		ctx.Call("fillText", fmt.Sprintf("%d", i), 10, 160-i+6)
	}
	for i := 0; i < 5; i++ {
		ctx.Call("fillText", fmt.Sprintf("%d", i+2018), i*40+25, 180)
	}
	ctx.Call("restore")
}

func main() {
	title := "Visual variable"
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
	h2.Set("innerHTML", title)
	body.Call("appendChild", h2)

	canvas := doc.Call("createElement", "canvas")
	ctx := canvas.Call("getContext", "2d")

	// flipping vertically
	ctx.Call("scale", 1, -1)

	canvas.Set("width", size.w*scale)
	canvas.Set("height", size.h*scale)
	body.Call("appendChild", canvas)

	canvas.Get("style").Set("width", size.w)
	canvas.Get("style").Set("height", size.h)
	ctx.Call("scale", scale, scale)

	ctx.Set("font", "16px Roboto")
	ctx.Call("fillText", "line chart", 20, 20)
	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}

	data := []MyData{
		{"A", []int{90, 75, 72, 65, 60}},
		{"B", []int{30, 43, 60, 66, 70}},
	}

	mycanvas.drawLineChart(data)

	<-make(chan struct{})
}
