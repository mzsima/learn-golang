package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	R := make([]int, n) 
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &R[i])
	}

	var minv, maxv int
	maxv = -10e9
	minv = R[0]
	for i := 1; i < n; i++ {
		maxv = Max(maxv, R[i] - minv)
		minv = Min(minv, R[i])
	}

	fmt.Printf("%d\n", maxv)
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}