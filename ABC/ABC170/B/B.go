package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y);

	for i:=0; i<=x; i++ {
		a, b := i, x - i
		if (2 * a + 4 * b == y) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}