package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("numbers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for id := range [10000]int{} {
		if id != 0 {
			w.WriteString(",")
		}
		w.WriteString(strconv.Itoa(id))
	}
	w.Flush()
}