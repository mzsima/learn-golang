package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var a [200]int

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	even := true
	count := 0
	for even {
		for i := 0; i < n; i++ {
			if a[i] % 2 != 0 { 
				even = false
			}
			a[i] /= 2
		}
		if even {
			count += 1
		}
	}

	fmt.Printf("%d\n", count)
}