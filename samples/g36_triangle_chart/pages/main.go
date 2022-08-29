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
	value float64
}

func (m MyCanvas) drawTriangle(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 4)
	ctx.Set("strokeStyle", "green")
	ctx.Call("translate", 80, 220)

	r := data[0].value
	pts := []Position{
		{r * 0.5, -r * math.Sin(math.Pi/3.0)},
		{r, 0},
		{0., 0.},
	}

	ctx.Call("beginPath")
	ctx.Call("moveTo", 0, 0)
	for i := 0; i < 3; i++ {
		ctx.Call("lineTo", pts[i].x, pts[i].y)
		ctx.Call("stroke")
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

	canvas.Set("width", size.w*scale)
	canvas.Set("height", size.h*scale)
	body.Call("appendChild", canvas)

	canvas.Get("style").Set("width", size.w)
	canvas.Get("style").Set("height", size.h)
	ctx.Call("scale", scale, scale)

	ctx.Set("font", "16px Roboto")
	ctx.Call("fillText", "triangle chart", 20, 20)
	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}

	data := []MyData{{200.}}

	mycanvas.drawTriangle(data)

	<-make(chan struct{})
}
