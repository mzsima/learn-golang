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
	m := readInt()

	rooms := make([]map[int]interface{}, n+1)
	for i := 0; i < n+1; i++ {
		rooms[i] = make(map[int]interface{})
	}
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		rooms[a][b] = true
		rooms[b][a] = true
	}

	HUGE := int(1e8)
	d := make([]int, n+1)
	for i := range d {
		d[i] = HUGE
	}
	d[1] = 0
	queue := []int{1}
	for len(queue) > 0 {
		roomNo := queue[0]
		queue = queue[1:]
		routes := rooms[roomNo]
		for r := range routes {
			if d[r] == HUGE {
				d[r] = d[roomNo] + 1
				queue = append(queue, r)
			}
		}
	}

	if d[n] == HUGE {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
		for roomNo := 2; roomNo < n+1; roomNo++ {
			min := n + 1
			next := roomNo
			for i := range rooms[roomNo] {
				if d[i] < min {
					min = d[i]
					next = i
				}
			}
			fmt.Println(next)
		}
	}

}
