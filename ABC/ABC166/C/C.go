package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
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
	defer w.Flush()

	n := readInt()
	m := readInt()
	tw := make([]int, n)
	for i := 0; i < n; i++ {
		tw[i] = readInt()
	}

	good := make([]bool, n)
	for i := 0; i < n; i++ {
		good[i] = true
	}

	for i := 0; i < m; i++ {
		a := readInt() - 1
		b := readInt() - 1
		if tw[a] > tw[b] {
			good[b] = false
		} else if tw[a] == tw[b] {
			good[a] = false
			good[b] = false
		} else {
			good[a] = false
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if good[i] {
			cnt++
		}
	}

	fmt.Println(cnt)
}
