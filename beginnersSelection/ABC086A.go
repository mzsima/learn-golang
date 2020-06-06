package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	ans := "Even"
	if a * b % 2 != 0 { ans = "Odd" }
	fmt.Printf("%s\n", ans)
}