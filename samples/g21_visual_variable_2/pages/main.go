package main

import (
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

func (m MyCanvas) drawGrayScale() {
	ctx := m.ctx
	ctx.Call("save")
	gradient := ctx.Call("createLinearGradient", 20, 0, 400, 0)
	gradient.Call("addColorStop", 0, "black")
	gradient.Call("addColorStop", 1, "white")
	ctx.Set("fillStyle", gradient)
	ctx.Call("beginPath")
	ctx.Call("rect", 20, 40, 400, 40)
	ctx.Call("fill")
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
	ctx.Call("fillText", "gray scale", 20, 20)

	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}
	mycanvas.drawGrayScale()

	<-make(chan struct{})
}
