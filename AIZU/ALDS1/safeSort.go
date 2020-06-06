package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	mark  string
	value int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	A := make([]Card, n)
	B := make([]Card, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		num, _ := strconv.Atoi(string([]rune(s)[1]))
		A[i] = Card{string([]rune(s)[0]), num}
		B[i] = Card{string([]rune(s)[0]), num}
	}

	Bubbled := BubbleSort(A, n)
	Selectioned := SelectionSort(B, n)

	BubbledStr := display(Bubbled)
	fmt.Println(BubbledStr)
	fmt.Println("Stable")

	SelectionedStr := display(Selectioned)
	fmt.Println(SelectionedStr)

	if BubbledStr == SelectionedStr {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func BubbleSort(C []Card, N int) []Card {
	for i := 0; i < N; i++ {
		for j := N - 1; j > i; j-- {
			if C[j].value < C[j-1].value {
				C[j], C[j-1] = C[j-1], C[j]
			}
		}
	}
	return C
}

func SelectionSort(C []Card, N int) []Card {
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if C[j].value < C[minj].value {
				minj = j
			}
		}
		C[i], C[minj] = C[minj], C[i]
	}
	return C
}

func display(cards []Card) string {
	var s []string
	for _, c := range cards {
		s = append(s, fmt.Sprintf("%s%d", c.mark, c.value))
	}
	return strings.Join(s, " ")
}
