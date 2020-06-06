package main

import (
	"fmt"
	"bufio"
	"os"
)

type testimony struct {
	x int
	y int
}

func main() {

	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)

	data := make([][]testimony, n)
	for i := 0; i < n; i++ {
		var a, x, y int
		fmt.Fscan(r, &a)

		t := make([]testimony, a)
		for j := 0; j < a; j++ {
			fmt.Fscan(r, &x)
			fmt.Fscan(r, &y)
			t[j] = testimony{x, y}
		}
		data[i] = t
	}
	
	ans := 0
	for i := 0; i < (1 << uint(n)); i++ {
		ok := true
		for j := 0; j < n; j++ {
			honest := i >> uint(j) & 1 == 1
			if !honest {
				continue
			}
			for _, t := range data[j] {
				if i >> uint(t.x - 1) & 1 != t.y {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}

		if ok {
			count := 0
			for j := 0; j < i; j++ {
				count += i >> uint(j) & 1
			}
			ans = Max(ans, count)
		}
	}

	fmt.Println(ans)
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}