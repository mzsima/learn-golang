package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	http.Handle("/", http.FileServer(http.Dir("pages/")))

	http.HandleFunc("/kline", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(getKline())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	log.Println(http.ListenAndServe(":8080", nil))

}

func getKline() *CandleStickResponse {
	resp, err := http.Get("https://api.coin.z.com/public/v1/klines?symbol=BTC&interval=30min&date=20220407")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jsonData CandleStickResponse
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	return &jsonData
}
