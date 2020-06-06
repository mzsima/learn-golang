package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func readInt() (read int) {
	scanner.Scan()
	read, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func dfs(s []string, a int, b int, c int) []string {
	if a < 0 || b < 0 || c < 0 {
		return nil
	}
	if len(s) == 0 {
		return []string{}
	}

	if "AB" == s[0] {
		if x := dfs(s[1:], a+1, b-1, c); x != nil {
			return append(x, "A")
		}
		if x := dfs(s[1:], a-1, b+1, c); x != nil {
			return append(x, "B")
		}
	} else if "AC" == s[0] {
		if x := dfs(s[1:], a+1, b, c-1); x != nil {
			return append(x, "A")
		}
		if x := dfs(s[1:], a-1, b, c+1); x != nil {
			return append(x, "C")
		}
	} else if "BC" == s[0] {
		if x := dfs(s[1:], a, b+1, c-1); x != nil {
			return append(x, "B")
		}
		if x := dfs(s[1:], a, b-1, c+1); x != nil {
			return append(x, "C")
		}
	}
	return nil
}

func main() {
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

	n := readInt()
	a := readInt()
	b := readInt()
	c := readInt()

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readString()
	}

	ans := dfs(s, a, b, c)

	if len(ans) > 0 {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
		return
	}

	for i := len(ans) - 1; i > -1; i-- {
		fmt.Fprintln(w, ans[i])
	}
}
