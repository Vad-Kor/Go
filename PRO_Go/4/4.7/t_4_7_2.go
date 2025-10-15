package main

import (
	"fmt"
)

func main_472() {
	var a, b, k int

	fmt.Scan(&a, &b, &k)

	for i := a; i <= b; i++ {
		iCount := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				iCount++
			}
		}
		if iCount >= k {
			if i != a {
				fmt.Print(" ")
			}
			fmt.Print(i)
		}
	}
}
