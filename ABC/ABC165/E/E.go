package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n%2 == 1 {
		for i := 1; i <= m; i++ {
			fmt.Printf("%d %d\n", i, n-i)
		}
	} else {
		flag := true
		for l, r := 1, n-1; l < r && l <= m; l, r = l+1, r-1 {
			if flag && r-l <= n/2 {
				flag = false
				r--
			}
			fmt.Printf("%d %d\n", l, r)
		}
	}
}
