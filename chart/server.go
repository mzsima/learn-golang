package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("." + r.RequestURI)
		if err == nil {
			io.Copy(w, file)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))

}
