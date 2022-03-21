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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

const MaxBuf = 200100

var buf []byte = make([]byte, MaxBuf)

type Point struct{ x, y int }

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(buf, MaxBuf)
	defer w.Flush()

	n := readInt()
	pts := make([]Point, n)

	for i := 0; i < n; i++ {
		x := readInt()
		y := readInt()
		pts[i] = Point{x, y}
	}

	s := readString()
	left_max := make(map[int]int)
	right_min := make(map[int]int)

	for i, pt := range pts {
		if string(s[i]) == "R" {
			if val, ok := left_max[pt.y]; ok && pt.x < val {
				fmt.Print("Yes")
				return
			}
		} else {
			if val, ok := right_min[pt.y]; ok && pt.x > val {
				fmt.Print("Yes")
				return
			}
		}

		if string(s[i]) == "R" {
			if _, ok := right_min[pt.y]; ok {
				right_min[pt.y] = minInt(right_min[pt.y], pt.x)
			} else {
				right_min[pt.y] = pt.x
			}
		} else {
			if _, ok := left_max[pt.y]; ok {
				left_max[pt.y] = maxInt(left_max[pt.y], pt.x)
			} else {
				left_max[pt.y] = pt.x
			}
		}
	}

	fmt.Print("No")
}
