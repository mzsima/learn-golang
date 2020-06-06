package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	A0, count := selectionSort(A, n)

	fmt.Fprintln(w, strings.Join(Itos(A0), " "))
	fmt.Fprintln(w, count)
}

func selectionSort(A []int, n int) ([]int, int) {
	count := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if i != minj {
			A[i], A[minj] = A[minj], A[i]
			count += 1
		}
	}
	return A, count
}

func Itos(A []int) []string {
	var s []string
	for _, k := range A {
		s = append(s, strconv.Itoa(k))
	}
	return s
}
