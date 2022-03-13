package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func readInt() (read int) {
	scanner.Scan()
	read, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	scanner.Split(bufio.ScanWords)

	MOD := 998244353
	n := readInt()
	s := readInt()
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
	}

	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = readInt()
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < s+1; j++ {
			dp[i+1][j] += (2 * dp[i][j] % MOD)
			dp[i+1][j] %= MOD
			if j+A[i] <= s {
				dp[i+1][j+A[i]] += (dp[i][j] % MOD)
				dp[i+1][j+A[i]] %= MOD
			}
		}
	}

	fmt.Println(dp[n][s])
}
