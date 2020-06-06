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
	k := readInt()

	okasi := make([]int, n)

	for i := 0; i < k; i++ {
		d := readInt()
		for j := 0; j < d; j++ {
			a := readInt()
			okasi[a-1] = 1
		}
	}
	sum := 0
	for _, i := range okasi {
		sum += i
	}

	fmt.Fprintln(w, n-sum)
}
