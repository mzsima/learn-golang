package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	p int
	s int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()

	a := []int{}
	b := []Pair{}

	area := 0
	for i, v := range in {
		if v == '\\' {
			a = append(a, i)
		} else if v == '/' {
			if len(a) < 1 {
				continue
			}
			var j int
			a, j = a[:len(a)-1], a[len(a)-1]
			s := i - j
			area += s
			for len(b) > 0 && b[len(b)-1].p > j {
				s += b[len(b)-1].s
				b = b[:len(b)-1]
			}
			b = append(b, Pair{j, s})
		}
	}
	fmt.Println(area)
	fmt.Printf("%d", len(b))
	for _, p := range b {
		fmt.Printf(" %d", p.s)
	}
	fmt.Println()
}
