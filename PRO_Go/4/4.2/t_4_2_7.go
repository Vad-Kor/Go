package main

import (
	"fmt"
)

func main_427() {
	var n, x int
	fmt.Scan(&n)

	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x%2 == 0 && x%3 != 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
