package main

import (
	"fmt"
	"math"
)

func powInt(a int, b int) (res int) {
	res = int(math.Pow(float64(a), float64(b)))
	return
}

func main() {
	var x int
	fmt.Scan(&x)

	max := 0
	for i := 0; true; i++ {
		if powInt(i, 5)-powInt(i-1, 5) > 1e9 {
			max = i
			break
		}
	}

	for a := 0; a < max; a++ {
		for b := 0; b < max; b++ {
			if powInt(a, 5)-powInt(b, 5) == x {
				fmt.Printf("%d %d\n", a, b)
				return
			}
			if powInt(a, 5)-powInt(-b, 5) == x {
				fmt.Printf("%d %d\n", a, -b)
				return
			}
		}
	}
}
