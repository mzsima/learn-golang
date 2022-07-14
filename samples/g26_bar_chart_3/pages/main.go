package main

import (
	"fmt"
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
	price []int
}

func (m MyCanvas) drawBarChart(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Call("translate", 0, 300)
	ctx.Call("scale", 1, -1) // origin to lower left corner
	for i, v := range data {
		ctx.Call("save")
		for j := 0; j < 3; j++ {
			ctx.Set("fillStyle", fmt.Sprintf("hsl(40, 0%%, %d%%)", j*20+30))
			ctx.Call("beginPath")
			h := 0
			for _, v := range v.price[:j] {
				h += v
			}
			ctx.Call("rect", i*70+20, h, 60, v.price[j])
			ctx.Call("fill")
		}
		ctx.Call("restore")
	}
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Set("font", "12px Roboto")
	for i := 0; i <= 250; i += 50 {
		ctx.Call("fillText", fmt.Sprintf("%d点", i), 70*len(data)+20, 300-i+6)
	}
	ctx.Call("restore")

	for i, v := range data {
		ctx.Call("fillText", v.label, i*70+35, 320)
	}
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
	ctx.Call("fillText", "bar chart (点数)", 20, 20)
	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}

	data := []MyData{
		{"A君", []int{100, 86, 82}},
		{"B君", []int{90, 74, 86}},
		{"C君", []int{50, 70, 60}},
	}

	mycanvas.drawBarChart(data)

	<-make(chan struct{})
}
