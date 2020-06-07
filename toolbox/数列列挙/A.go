package main

import "fmt"

var n, m int

func dfs(A []int) {
	if len(A) == n+1 {
		fmt.Println(A[1:])
		return
	}

	for l := 0; l <= m; l++ {
		dfs(append(A, l))
	}
}

func main() {
	n = 3
	m = 9
	dfs([]int{1})
}
