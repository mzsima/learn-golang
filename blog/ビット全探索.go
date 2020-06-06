package main

import (
	"fmt"
)

func main() {
	n := 3
	for i := 0; i < (1 << uint(n)); i++ {
		fmt.Printf("%03b\n", i)
		if i >> uint(1) & 1 == 1{
			// ビットの ２桁目が立っている場合 
			fmt.Println("hit")
		} 
	}
}