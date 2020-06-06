package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500 * i + 100 * j + 50 * k == x { count += 1 }
			} 
		} 
	}

	fmt.Printf("%d\n", count) 
}