package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
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

	for i := 1; i<=5; i++ {
		x := readInt()
		if x == 0 {
			fmt.Fprintln(w, i)
			return
		}
	}
}