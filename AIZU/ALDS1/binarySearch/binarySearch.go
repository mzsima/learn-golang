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
	S := make([]int, n)
	for i := 0; i < n; i++ {
		S[i] = readInt()
	}

	cnt := 0
	q := readInt()
	for j := 0; j < q; j++ {
		t := readInt()
		if binarySearch(S, t, n) > -1 {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func binarySearch(A []int, key int, n int) int {
	left := 0
	right := n
	for left < right {
		mid := (left + right) / 2
		if A[mid] == key {
			return mid
		} else if key < A[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
