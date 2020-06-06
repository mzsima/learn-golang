package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	if x < 2 || x%2 == 0 {
		return false
	}

	for i := 3; float64(i) <= math.Sqrt(float64(x)); i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	if isPrime(n) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
