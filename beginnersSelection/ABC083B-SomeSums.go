package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d%d%d", &n, &a, &b)
	
	sum := 0
	for i := 1; i <= n; i++ {
		v := calc(i)
		if v >= a && v <= b {
			sum += i
		}
	}

	fmt.Printf("%d\n", sum)
}

func calc(x int) int {
	val := 0
	for x > 0 {
		val += x % 10
		x /= 10 
	}
	return val
}