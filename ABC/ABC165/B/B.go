package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	b := 100
	for i := 1; true; i++ {
		b = int(float64(b) * 1.01)
		if b >= x {
			fmt.Println(i)
			return
		}
	}
}
