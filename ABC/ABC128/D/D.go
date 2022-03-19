package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func maxInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[len(a) - 1]
}

func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	n := readInt()
	k := readInt()
	v := []int{}
	for i:=0; i<n; i++ {
		v = append(v, readInt())
	}

	nk := minInt([]int{n, k})
	max := 0
	for i:=0; i<=nk; i++ {
		for j:=0; j<=nk-i; j++ {
			left := v[:i]
			right := v[len(v) - j:]
			elm := append([]int{}, left...)
			elm = append(elm, right...)
			sum := 0
			negatives := []int{}

			for _, p := range elm {
				if p < 0 { 
					negatives = append(negatives, p) 
				}
				sum += p
			}

			delcnt := minInt([]int{k - i - j, len(negatives)})
			sort.Sort(sort.IntSlice(negatives))
			for x:=0; x<delcnt; x++ {
				sum -= negatives[0]
				negatives = negatives[1:]
			}
			if max < sum {
				max = sum
			}
		}
	}
	fmt.Print(max)
}