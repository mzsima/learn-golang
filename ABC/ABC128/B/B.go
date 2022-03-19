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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

type Restaurant struct {
	id int
	name string
	point int
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	n := readInt()
	s := make([]Restaurant, n)

	for i:=0; i<n; i++ {
		name := readString()
		point := readInt()
		s[i] = Restaurant{i+1, name, point}
	}
	
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].name == s[j].name {
			 return s[i].point > s[j].point
		}
		return s[i].name < s[j].name
	})

	for _, r := range s {
		fmt.Fprintln(w, r.id)
	}
}