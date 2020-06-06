package main

func mulOverflows(a, b int) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func main() {

}
