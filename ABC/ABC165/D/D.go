package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	x := 0

	if b > n {
		x = n
	} else {
		x = b - 1
	}
	fmt.Println(int(float64(a)*float64(x)/float64(b)) - a*int(float64(x)/float64(b)))
}
