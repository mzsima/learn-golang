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

type Feeling struct {
	Name string
}

type Face struct {
	canvas js.Value
	ctx    js.Value
	center Position
}

func (f Face) drawOuter() {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y+100)
	ctx.Set("strokeStyle", "gray")
	ctx.Set("lineWidth", 10)

	ctx.Call("save")
	ctx.Call("beginPath")
	ctx.Call("ellipse", 0, 0, 100, 100, 0, 0, math.Pi*2.0)
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) drawTop() {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("fillStyle", "gray")

	ctx.Call("save")
	ctx.Call("translate", -60, 20)
	ctx.Call("fillRect", -50, -5, 100, 10)
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Call("translate", 60, 20)
	ctx.Call("fillRect", -50, -5, 100, 10)
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) drawMiddle() {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("strokeStyle", "gray")
	ctx.Set("lineWidth", 10)

	ctx.Call("save")
	ctx.Call("translate", -60, 80)
	ctx.Call("beginPath")
	ctx.Call("arc", 0, 0, 24, 0, 2*math.Pi)
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Call("translate", 60, 80)
	ctx.Call("beginPath")
	ctx.Call("arc", 0, 0, 24, 0, 2*math.Pi)
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) drawBottom() {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("strokeStyle", "gray")
	ctx.Set("lineWidth", 10)

	ctx.Call("save")
	ctx.Call("translate", 0, 200)
	ctx.Call("beginPath")
	ctx.Call("ellipse", 0, 0, 64, 64, 0, math.Pi*1.2, math.Pi*1.8)
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) draw() {
	f.drawOuter()
	f.drawTop()
	f.drawMiddle()
	f.drawBottom()
}

func main() {
	title := "Chernoff face"
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

	face := Face{canvas, ctx, Position{150, 30}}
	face.draw()

	<-make(chan struct{})
}
