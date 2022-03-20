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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

type Point struct{ x, y int }

const MaxBuf = 200100

var buf []byte = make([]byte, MaxBuf)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(buf, MaxBuf)
	defer w.Flush()

	v := readInt()
	a := readInt()
	b := readInt()
	c := readInt()

	for i := 0; i < 100000; i++ {
		for i, u := range []int{a, b, c} {
			v -= u
			if v < 0 {
				switch i {
				case 0:
					fmt.Print("F")
				case 1:
					fmt.Print("M")
				case 2:
					fmt.Print("T")
				}
				return
			}
		}
	}
}
