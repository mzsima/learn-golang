package main

import (
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
	value int
	color string
}

func (m MyCanvas) drawPieChart(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Call("translate", 160, 160)

	total := 0
	for _, d := range data {
		total += d.value
	}

	s := 0.
	for _, d := range data {
		v := float64(d.value) / float64(total)
		ctx.Set("fillStyle", d.color)
		ctx.Call("beginPath")
		ctx.Call("moveTo", 0, 0)
		ctx.Call("arc", 0, 0, 120, s*math.Pi*2.0, (s+v)*math.Pi*2.0)
		ctx.Call("lineTo", 0, 0)
		ctx.Call("fill")
		s += v
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
	ctx.Call("fillText", "pie chart", 20, 20)
	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}

	data := []MyData{{80, "#5A0273"}, {20, "#078C03"}, {200, "#F29F05"}, {105, "#F25C05"}, {38, "#A60303"}}

	mycanvas.drawPieChart(data)

	<-make(chan struct{})
}
