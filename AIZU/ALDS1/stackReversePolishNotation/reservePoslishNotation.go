package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in []string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in = strings.Split(scanner.Text(), " ")
	}

	var n1, n2 int
	stack := []int{}
	for _, s := range in {
		switch s {
		case "+":
			if len(stack) > 1 {
				n1, n2, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
				stack = append(stack, n1+n2)
			}
		case "-":
			if len(stack) > 1 {
				n1, n2, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
				stack = append(stack, n1-n2)
			} else if len(stack) == 0 {
				stack[0] = stack[0] * -1
			}
		case "*":
			n1, n2, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
			stack = append(stack, n1*n2)
		default:
			i, _ := strconv.Atoi(s)
			stack = append(stack, i)
		}
	}
	fmt.Println(stack[0])
}
