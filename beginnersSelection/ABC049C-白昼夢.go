package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	r := reverse(s)	
	words := []string{ reverse("dream"), reverse("dreamer"), reverse("erase"), reverse("eraser") }

	hit := true
	for hit {
		hit = false
		for _, w := range words {		
			if strings.HasPrefix(r, w) {
				hit = true
				r = strings.TrimPrefix(r, w)
				break
			}
		}
	}

	if len(r) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}