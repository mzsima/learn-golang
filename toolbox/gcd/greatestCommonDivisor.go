package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d%d", &x, &y)
	ans := gcd(x, y)
	fmt.Printf("%d\n", ans)
}

func gcd(x int, y int) int {
	var a, b int
	if x > y {
		a, b = x, y
	} else {
		a, b = y, x
	}

	for b > 0 {
		r := a % b
		a = b
		b = r
	}

	return a
}