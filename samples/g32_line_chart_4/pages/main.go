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
	label string
	value int
}

func (m MyCanvas) drawLineChart(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Call("translate", 0, 160)
	ctx.Call("scale", 1, -1) // origin to lower left corner

	ctx.Call("beginPath")
	for i, d := range data {
		x := i*60 + 40
		ctx.Call("lineTo", x, d.value)
		ctx.Call("stroke")
		ctx.Call("beginPath")
		ctx.Call("arc", x, d.value, 3, 0, math.Pi*2.0)
		ctx.Call("fill")
	}
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Set("font", "12px Roboto")
	for i := 0; i <= 100; i += 20 {
		ctx.Call("fillText", fmt.Sprintf("%d", i), 10, 160-i+6)
	}
	for i, v := range data {
		ctx.Call("fillText", fmt.Sprintf("%s", v.label), i*60+25, 180)
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
		{"2017", 30},
		{"2018", 30},
		{"2019", 60},
		{"2020", 90},
	}

	mycanvas.drawLineChart(data)

	ctx.Call("translate", 300, 0)
	mycanvas.drawLineChart(append(data[:1], data[2:]...))

	<-make(chan struct{})
}
