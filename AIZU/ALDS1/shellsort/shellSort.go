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

	shellSort(A, n)
}

func insertionSort(A []int, n int, g int) int {
	cnt := 0
	for i := g; i < n; i++ {
		v := A[i]
		j := i - g
		for j >= 0 && A[j] > v {
			A[j+g] = A[j]
			j = j - g
			cnt++
		}
		A[j+g] = v
	}
	return cnt
}

func shellSort(A []int, n int) {
	cnt := 0

	G := []int{}
	for i := 1; i <= n; i = 3*i + 1 {
		G = append(G, i)
	}
	m := len(G)
	fmt.Println(m)

	Gstr := []string{}
	for i := len(G) - 1; i >= 0; i-- {
		Gstr = append(Gstr, strconv.Itoa(G[i]))
	}
	fmt.Println(strings.Join(Gstr, " "))

	for i := 0; i < m; i++ {
		cnt += insertionSort(A, n, G[m-i-1])
	}
	fmt.Println(cnt)

	for _, a := range A {
		fmt.Println(a)
	}

}
