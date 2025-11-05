package main

import (
	"fmt"
)

func main_4107() {
	var n int
	fmt.Scan(&n)
	max := 0
	maxPrev := 0

	for n != 0 {
		if n > max {
			maxPrev = max
			max = n
		} else if n > maxPrev {
			maxPrev = n
		}
		fmt.Scan(&n)
	}
	fmt.Println(maxPrev)
}
