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

const M = 10000000

type Dict struct {
	dict []bool
}

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

func getChar(ch rune) int {
	switch ch {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	}
	// 'T'
	return 4
}

func getKey(s string) int {
	keys := []rune(s)
	return h2(h1(keys))
}

func h1(keys []rune) int {
	sum := 0
	p := 1
	for i := 0; i < len(keys); i++ {
		sum += p * (getChar(keys[i]))
		p *= 5
	}
	return sum
}

func h2(key int) int {
	return key % M
}

func (d *Dict) find(s string) bool {
	key := getKey(s)
	return d.dict[key] == true
}

func (d *Dict) insert(s string) {
	key := getKey(s)
	d.dict[key] = true
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	n := readInt()

	d := Dict{make([]bool, M)}

	for i := 0; i < n; i++ {
		cmd := readString()
		v := readString()
		if cmd == "insert" {
			d.insert(v)
		} else {
			if d.find(v) {
				fmt.Fprintln(w, "yes")
			} else {
				fmt.Fprintln(w, "no")
			}
		}
	}
}
