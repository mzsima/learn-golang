package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)

	for i := 0; i < n ; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	totalA := 0
	totalB := 0
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			totalA += a[n - i - 1]
		} else {
			totalB += a[n - i - 1]
		}
	}

	fmt.Printf("%d\n", totalA - totalB)
}