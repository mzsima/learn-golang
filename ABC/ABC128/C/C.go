package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	l := make([][]int, m)
	p := make([]int, m)
	for i:=0; i<m; i++ {
		var t int
		fmt.Scan(&t)
		l[i] = make([]int, t)
		for j:=0; j<t; j++ {
			fmt.Scan(&l[i][j])
		}
	}
	for i:=0; i<m; i++ {
		fmt.Scan(&p[i])
	}

	allOn := 0
	for bit:=0; bit<(1<<uint(n)); bit++{
		check := true
		for i:=0; i<m; i++ {
			cnt := 0
			for _, v := range l[i] {
				if (bit>>uint(v-1))&1 == 1 {
					cnt++
				}
			}
			if p[i] != (cnt % 2) {
				check = false
			}
		}
		if check {
			allOn++
		}
	}
	fmt.Print(allOn)
}