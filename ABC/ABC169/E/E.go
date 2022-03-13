package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	mins := make([]int, n)
	maxs := make([]int, n)

	for i := 0; i < n; i++ {
		mins[i] = readInt()
		maxs[i] = readInt()
	}

	sort.Ints(mins)
	sort.Ints(maxs)

	cnt := 0
	if n%2 == 0 {
		c1 := mins[n/2-1] + mins[n/2]
		c2 := maxs[n/2-1] + maxs[n/2]
		cnt = c2 - c1 + 1
	} else {
		c1 := mins[n/2]
		c2 := maxs[n/2]
		cnt = c2 - c1 + 1
	}
	fmt.Println(cnt)
}
