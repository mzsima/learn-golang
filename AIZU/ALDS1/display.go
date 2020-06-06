package main

import (
	"fmt"
	"strconv"
	"strings"
)

func display(A []int) {
	var displays []string
	for _, k := range A {
		displays = append(displays, strconv.Itoa(k))
	}
	fmt.Println(strings.Join(displays, " "))
}
