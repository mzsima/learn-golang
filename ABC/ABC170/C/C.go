package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	w		= bufio.NewWriter(os.Stdout)
)

func readInt() (read int) {
	scanner.Scan()
	read, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return
}

func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	x := readInt()
	y := readInt()

	if y == 0 {
		fmt.Fprintln(w, x)
		return
	}

	nums := make(map[int]bool)
	for i := 0; i<y; i++ {
		p := readInt()
		nums[p] = true
	}
	
	for i:=0; i <= x+1; i++ {
		ap, am := x + i, x - i;

		if _, ok := nums[am]; !ok {
			fmt.Fprintln(w, am)
			return
		}
		if _, ok := nums[ap]; !ok {
			fmt.Fprintln(w, ap)
			return
		}
	}

	
}