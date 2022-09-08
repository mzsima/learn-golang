package main

import (
	"math"
	"syscall/js"
)

type Position struct {
	x float64
	y float64
}

type Vector struct {
	x float64
	y float64
}

func (v *Vector) mul(s float64) {
	v.x *= s
	v.y *= s
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
	l  float64
	v1 float64
	v2 float64
	v3 float64
}

func (m MyCanvas) drawTriangle(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Set("strokeStyle", "green")
	ctx.Call("translate", 80, 220)

	r := data[0].l
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

func (m MyCanvas) drawVectors(data []MyData) {
	ctx := m.ctx

	ctx.Call("save")
	ctx.Set("lineWidth", 2)
	ctx.Set("strokeStyle", "green")
	ctx.Call("translate", 80, 220)

	r := data[0].l
	pt := []Position{
		{r * 0.5, -r * math.Sin(math.Pi/3.0)},
		{r, 0},
		{0., 0.},
	}
	vectors := []Vector{
		{pt[1].x - pt[0].x, pt[1].y - pt[0].y},
		{pt[2].x - pt[1].x, pt[2].y - pt[1].y},
		{pt[0].x - pt[2].x, pt[0].y - pt[2].y},
	}
	axis := []Vector{
		{pt[1].x - pt[0].x, pt[1].y - pt[0].y},
		{pt[2].x - pt[1].x, pt[2].y - pt[1].y},
		{pt[0].x - pt[2].x, pt[0].y - pt[2].y},
	}
	vectors[0].mul(data[0].v1)
	vectors[1].mul(data[0].v2)
	vectors[2].mul(data[0].v3)

	for i := 0; i < 3; i++ {
		ctx.Call("beginPath")
		ctx.Call("moveTo", pt[i].x, pt[i].y)
		ctx.Set("lineWidth", 10)
		ctx.Call("lineTo", pt[i].x+vectors[i].x, pt[i].y+vectors[i].y)
		ctx.Call("stroke")
		ctx.Set("lineWidth", 1)
		ctx.Call("lineTo", pt[i].x+vectors[i].x-axis[(i+2)%3].x, pt[i].y+vectors[i].y-axis[(i+2)%3].y)
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

	data := []MyData{{200., 0.2, 0.5, 0.3}}

	mycanvas.drawTriangle(data)
	mycanvas.drawVectors(data)

	<-make(chan struct{})
}
