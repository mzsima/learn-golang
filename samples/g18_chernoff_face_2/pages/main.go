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

type TestScore struct {
	name    string
	math    uint
	science uint
	music   uint
}

func (t TestScore) toString() string {
	return fmt.Sprintf("氏名: %v, 数学: %v, 理科: %v, 音楽: %v", t.name, t.math, t.science, t.music)
}

func (t TestScore) fellOfSubject(subject string) Feeling {
	var score uint
	switch subject {
	case "math":
		score = t.math
	case "science":
		score = t.science
	case "music":
		score = t.music
	}

	if score < 40 {
		return Bad
	} else if score < 70 {
		return Good
	} else {
		return Excellent
	}
}

type Feeling string

type Face struct {
	canvas js.Value
	ctx    js.Value
	center Position
}

const (
	Excellent Feeling = "Better"
	Good      Feeling = "Good"
	Bad       Feeling = "Bad"
)

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

func (f Face) drawTop(feel Feeling) {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("fillStyle", "gray")

	ctx.Call("save")
	ctx.Call("translate", -60, 30)
	switch feel {
	case Excellent:
		ctx.Call("rotate", math.Pi*0.1)
	case Bad:
		ctx.Call("rotate", -math.Pi*0.1)
	}
	ctx.Call("fillRect", -50, -5, 100, 10)
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Call("translate", 60, 30)
	switch feel {
	case Excellent:
		ctx.Call("rotate", -math.Pi*0.1)
	case Bad:
		ctx.Call("rotate", math.Pi*0.1)
	}
	ctx.Call("fillRect", -50, -5, 100, 10)
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) drawMiddle(feel Feeling) {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("strokeStyle", "gray")
	ctx.Set("fillStyle", "gray")
	ctx.Set("lineWidth", 10)

	ctx.Call("save")
	ctx.Call("translate", -50, 80)
	ctx.Call("beginPath")
	switch feel {
	case Excellent:
		ctx.Call("arc", 0, 0, 24, 0, 2*math.Pi)
	case Good:
		ctx.Call("ellipse", 0, 0, 24, 18, 0, 0, 2*math.Pi)
	case Bad:
		ctx.Call("fillRect", -20, -5, 40, 10)
	}
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("save")
	ctx.Call("translate", 50, 80)
	ctx.Call("beginPath")
	switch feel {
	case Excellent:
		ctx.Call("arc", 0, 0, 24, 0, 2*math.Pi)
	case Good:
		ctx.Call("ellipse", 0, 0, 24, 18, 0, 0, 2*math.Pi)
	case Bad:
		ctx.Call("fillRect", -20, -5, 40, 10)
	}
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) drawBottom(feel Feeling) {
	ctx := f.ctx
	ctx.Call("save")
	ctx.Call("translate", f.center.x, f.center.y)
	ctx.Set("strokeStyle", "gray")
	ctx.Set("fillStyle", "gray")
	ctx.Set("lineWidth", 10)

	ctx.Call("save")
	ctx.Call("translate", 0, 200)
	ctx.Call("beginPath")
	switch feel {
	case Excellent:
		ctx.Call("ellipse", 0, -90, 64, 64, 0, math.Pi*0.2, math.Pi*0.8)
	case Good:
		ctx.Call("fillRect", -20, -50, 40, 10)
	case Bad:
		ctx.Call("ellipse", 0, 0, 64, 64, 0, math.Pi*1.2, math.Pi*1.8)
	}
	ctx.Call("stroke")
	ctx.Call("restore")

	ctx.Call("restore")
}

func (f Face) draw(score TestScore) {
	f.ctx.Call("save")
	f.ctx.Call("scale", 0.25, 0.25)
	f.drawOuter()
	f.drawTop(score.fellOfSubject("math"))
	f.drawMiddle(score.fellOfSubject("science"))
	f.drawBottom(score.fellOfSubject("music"))
	f.ctx.Call("restore")
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

	scores := []TestScore{
		{"Tanaka-san", 30, 39, 40},
		{"Sato-san", 80, 80, 80},
		{"Suzuki-san", 40, 80, 30},
		{"Yamamoto-san", 60, 60, 60},
	}

	for i, score := range scores {
		ctx.Call("fillText", score.toString(), 50, 100+20*i)
	}

	for i, score := range scores {
		face := Face{canvas, ctx, Position{250 * float64(i+1), 50}}
		face.draw(score)
	}
	<-make(chan struct{})
}
