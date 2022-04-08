package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

type CandleStickResponse struct {
	Status       int
	Responsetime string
	Data         []Candlestick
}

type Candlestick struct {
	OpenTime int64   `json:",string"`
	Open     int64   `json:",string"`
	High     int64   `json:",string"`
	Low      int64   `json:",string"`
	Close    int64   `json:",string"`
	Volume   float64 `json:",string"`
}

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
	h2.Set("innerHTML", "Chart01 WebAssembly! (made with GO)")
	body.Call("appendChild", h2)

	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", 500)
	canvas.Set("height", 500)
	body.Call("appendChild", canvas)

	ctx := canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "rgba(128, 128, 128, 0.5)")
	ctx.Call("beginPath")

	ctx.Call("translate", 40, 40)

	resp, err := http.Get("/kline")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var jsonData CandleStickResponse
	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	for i, v := range jsonData.Data {
		ctx.Call("fillRect", i*10-1, 1100-(v.High-5000000)/400.0, 2, (v.High-v.Low)/400.0)
		if v.Open > v.Close {
			ctx.Call("fillRect", i*10-3, 1100-(v.Open-5000000)/400.0, 6, (v.Open-v.Close)/400.0)
		} else {
			ctx.Call("fillRect", i*10-3, 1100-(v.Close-5000000)/400.0, 6, (v.Close-v.Open)/400.0)
		}
	}

}
