package main

import (
	"fmt"
)

func main_345() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= 2; i++ {
		if (n+i)%2 == 0 {
			fmt.Println(n + i)
			break
		}
	}
}
