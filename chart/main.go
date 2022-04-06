package main

import (
	"math"
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	if !doc.Truthy() {
		panic("unable to get 'document object")
	}

	body := doc.Get("body")
	if !body.Truthy() {
		panic("unable to get 'body' object")
	}

	h2 := doc.Call("createElement", "h2")
	h2.Set("innerHTML", "Hello from WebAssembly! (made with GO)")
	body.Call("appendChild", h2)

	canvas := doc.Call("createElement", "canvas")
	body.Call("appendChild", canvas)

	ctx := canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "rgba(128, 128, 128, 0.2)")
	ctx.Call("beginPath")

	ctx.Call("translate", 40, 40)
	for i := 0; i < 32; i++ {
		ctx.Call("rotate", float32(i)*math.Pi/16.0)
		ctx.Call("fillRect", 10, 10, 40, 40)
	}

}
