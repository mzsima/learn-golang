package main

import "fmt"

func main() {
	var s, t int 
	fmt.Scanf("%d %d", &s, &t);

	var cnt = 0;
	for a:=0; a<=100; a++ {
		for b:=0; b<=100-a; b++ {
			for c:=0; c<=100-a-b; c++ {
				if (a + b + c <= s) && (a * b * c <= t) {
					cnt++;
				}
			}		
		}	
	}
	fmt.Printf("%d", cnt);	
}