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
	label  string
	values []int
}

func (m MyCanvas) drawLineChart(data []MyData) {
	ctx := m.ctx
	colors := []string{"orange", "green"}
	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Call("translate", 0, 160)
	ctx.Call("scale", 1, -1) // origin to lower left corner

	stackval := []int{}
	for j, d := range data {
		for i, v := range d.values {
			if j == 0 {
				stackval = append(stackval, v)
			} else {
				stackval[i] += v
			}
		}
	}

	for j, d := range data {
		ctx.Call("beginPath")
		ctx.Call("moveTo", 40, 0)
		for i, v := range d.values {
			x := i*40 + 40
			ctx.Call("lineTo", x, stackval[i])
			stackval[i] -= v
		}
		ctx.Call("lineTo", len(d.values)*40, 0)
		ctx.Set("fillStyle", colors[j])
		ctx.Call("fill")
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
		{"A", []int{40, 35, 22, 40, 45}},
		{"B", []int{30, 43, 60, 46, 25}},
	}

	mycanvas.drawLineChart(data)

	<-make(chan struct{})
}
