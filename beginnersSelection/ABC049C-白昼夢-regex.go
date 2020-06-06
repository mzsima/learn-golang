package main

import (
	"fmt"
	"regexp"
)

func main() {
	divide := `^(dream|dreamer|erase|eraser)*$`
	fmt.Scanf("%s", &s)
	if regexp.MustCompile(divide).Match([]byte(s)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}