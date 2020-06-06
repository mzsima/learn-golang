package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func readInt() (read int) {
	scanner.Scan()
	read, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := readInt()
	m := readInt()
	x := readInt()

	CA := make([][]int, n)
	for i := range CA {
		CA[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			CA[i][j] = readInt()
		}
	}

	ans := 0
	for i := 0; i < (1 << uint(n)); i++ {
		price := 0
		learn := make([]int, m)
		for k := 0; k < n; k++ {
			if i>>uint(k)&1 == 0 {
				price += CA[k][0]
				for l := 1; l < m+1; l++ {
					learn[l-1] += CA[k][l]
				}
			}
		}
		update := true
		for _, v := range learn {
			if v < x {
				update = false
			}
		}
		if update && (ans > price || ans == 0) {
			ans = price
		}
	}

	if ans > 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}

}
