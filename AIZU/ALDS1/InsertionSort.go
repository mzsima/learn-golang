package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &A[i])
	}
	display(A)

	for i := 1; i < n; i++ {
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v

		display(A)
	}
}

func display(A []int) {
	var displays []string
	for _, k := range A {
		displays = append(displays, strconv.Itoa(k))
	}
	fmt.Println(strings.Join(displays, " "))
}