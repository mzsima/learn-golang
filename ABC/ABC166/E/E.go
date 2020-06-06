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
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = readInt()
	}

	xi := make(map[int]int)
	xj := make(map[int]int)
	for i := 0; i < n; i++ {
		xi[i+A[i]]++
		xj[i-A[i]]++
	}

	cnt := 0
	for k, v := range xi {
		cnt += v * xj[k]
	}
	fmt.Fprintln(w, cnt)
}
