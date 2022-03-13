package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func mulOverflows(a, b int) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func main() {
	var n int
	fmt.Scan(&n)

	sc.Split(bufio.ScanWords)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		v := nextInt()
		if v == 0 {
			fmt.Println(0)
			return
		}
		A[i] = v
	}

	a := 1
	for i := 0; i < n; i++ {
		if mulOverflows(a, A[i]) {
			fmt.Println(-1)
			return
		}
		a *= A[i]
		if a > 1e18 {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(a)
}
