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

	n := readInt()

	aa := make([]int, n)
	ba := make([]int, n)

	am := make(map[int]int)
	bm := make(map[int]int)
	cnt := 0
	for i := 0; i < n; i++ {
		aa[i] = readInt()
	}
	for i := 0; i < n; i++ {
		ba[i] = readInt()
	}

	for i := 0; i < n; i++ {
		sa := aa[i]
		sb := ba[i]
		if sa == sb {
			cnt++
		} else {
			am[sa] = 1
			bm[sb] = 1
		}
	}

	subCnt := 0
	for k, v := range am {
		if v > 0 && bm[k] > 0 {
			subCnt += 1
		}
	}

	fmt.Println(cnt)
	fmt.Println(subCnt)
}
