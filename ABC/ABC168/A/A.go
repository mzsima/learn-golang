package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	word := "hon"
	switch i := n % 10; i {
	case 0, 1, 6, 8:
		word = "pon"
	case 3:
		word = "bon"
	}

	fmt.Println(word)
}
