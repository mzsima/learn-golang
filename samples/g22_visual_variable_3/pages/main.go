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

func (m MyCanvas) drawColorScale() {
	ctx := m.ctx
	ctx.Call("save")
	for i := 0; i < 480; i++ {
		ctx.Set("fillStyle", fmt.Sprintf("hsl(%d, 100%%, 50%%)", i/2))
		ctx.Call("beginPath")
		ctx.Call("rect", 20+i, 40, 1, 40)
		ctx.Call("fill")
		ctx.Call("restore")
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
	ctx.Call("fillText", "hsl scale (0 to 240)", 20, 20)

	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}
	mycanvas.drawColorScale()

	<-make(chan struct{})
}
