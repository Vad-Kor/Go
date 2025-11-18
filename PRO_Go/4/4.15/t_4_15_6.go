package main

import (
	"fmt"
)

func main_4156() {
	var n, d int
	fmt.Scan(&n)
	min := 0
	count := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		if d < min {
			min = d
			count = 1
		} else if d == min {
			count++
		}
	}

	fmt.Println(count)
}
