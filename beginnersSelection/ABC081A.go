package main

import (
	"fmt"
)

func main() {
	var s int
	fmt.Scanf("%d", &s)

	count := 0
	for s > 0 {
		count += (s % 2)
		s /= 10
	}

	fmt.Printf("%d\n", count)
}