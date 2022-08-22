package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("YES")
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	log.Println(http.ListenAndServe(":80", nil))
}
