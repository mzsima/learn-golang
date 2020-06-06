package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	count := 0
	for _, num := range nums {
		if isPrime(num) {
			count += 1
		}
	}

	fmt.Printf("%d\n", count)
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	if x < 2 || x % 2 == 0 {
		return false
	}
	
	for i := 3; float64(i) <= math.Sqrt(float64(x)); i += 2 {
		if x % i == 0 {
			return false
		}
	}

	return  true
}