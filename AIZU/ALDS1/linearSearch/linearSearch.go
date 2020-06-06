package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

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
	S := make([]int, n)
	for i := 0; i < n; i++ {
		S[i] = readInt()
	}

	cnt := 0
	q := readInt()
	for i := 0; i < q; i++ {
		t := readInt()
		for _, s := range S {
			if s == t {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
