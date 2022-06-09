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

func (m MyCanvas) drawCircleSizedByRadius(value float64, p Position) {
	r := value * 10
	ctx := m.ctx
	ctx.Call("save")
	ctx.Set("fillStyle", "gray")
	ctx.Call("beginPath")
	ctx.Call("ellipse", p.x, p.y, r, r, 0, 0, math.Pi*2.0)
	ctx.Call("fill")
	ctx.Call("restore")
}

func (m MyCanvas) drawCircleSizedByArea(value float64, p Position) {
	r := math.Sqrt(value) * 10
	ctx := m.ctx
	ctx.Call("save")
	ctx.Set("fillStyle", "gray")
	ctx.Call("beginPath")
	ctx.Call("ellipse", p.x, p.y, r, r, 0, 0, math.Pi*2.0)
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

	scores := []float64{
		1., 2., 3.,
	}

	ctx.Set("font", "16px Roboto")
	ctx.Call("fillText", "視覚変数に、半径(上段) or 面積(下段)", 20, 20)
	ctx.Call("fillText", "値は [1, 2, 3]", 20, 48)

	mycanvas := MyCanvas{canvas, ctx, Position{size.w * 0.5, size.h * 0.5}}
	for i, score := range scores {
		mycanvas.drawCircleSizedByRadius(score, Position{100*float64(i) + 50, 100})
		mycanvas.drawCircleSizedByArea(score, Position{100*float64(i) + 50, 200})
	}

	<-make(chan struct{})
}
