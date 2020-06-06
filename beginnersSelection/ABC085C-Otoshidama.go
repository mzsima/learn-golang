package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scanf("%d%d", &n, &y)

	for i := 0; i <= n; i++ {
		for j:=0; j <= n-i; j++ {
			k := n - i - j
			if y / 1000 == 10 * i + 5 * j + k {
				fmt.Printf("%d %d %d\n", i, j, k)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}