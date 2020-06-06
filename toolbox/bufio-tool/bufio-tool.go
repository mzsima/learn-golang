package main

import (
	"bufio"
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

func main() {
	scanner.Split(bufio.ScanWords)
	defer w.Flush()

}
