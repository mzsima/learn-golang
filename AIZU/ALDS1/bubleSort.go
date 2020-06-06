package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &A[i])
	}

	A0, count := bubbleSort(A, n)

	fmt.Fprintln(w, strings.Join(itoS(A0)," "))
	fmt.Fprintln(w, count)
}

func bubbleSort(A []int, n int) ([]int, int) {
	flag := true
	count := 0
	for flag {
		flag = false
		for j := n - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				buf := A[j]
				A[j] = A[j-1]
				A[j-1] = buf
				flag = true
				count += 1
			}
		}
	}
	return A, count
}

func itoS(A []int) []string {
	var s []string
	for _, k := range A {
		s = append(s, strconv.Itoa(k))
	}
	return s
}