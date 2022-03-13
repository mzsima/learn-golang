package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n);

	l := make([][][2]int, n);
	for i:=0; i<n; i++ {
		var a int
		fmt.Scan(&a)
		ansArr := make([][2]int, a)
		for j:=0; j<a; j++ {
			fmt.Scan(&ansArr[j][0], &ansArr[j][1])
		}
		l[i] = ansArr
	}

	max := 0
	for bit:=0; bit<(1<<uint(n)); bit++ {
		honest := 0
		for i:=0;i<n;i++ {
			if (bit>>uint(i))&1 == 1 {
				for _, ans := range l[i] {
					if bit>>uint(ans[0]-1)&1 != ans[1] {
						honest = -10000
						break;
					}
				}
				honest++
			}
		}
		if max < honest {
			max = honest
		}
	}

	fmt.Print(max)
}