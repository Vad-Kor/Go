package main

import (
	"fmt"
)

func main_464() {
	var n, n_prev int

	fmt.Scan(&n)
	count := 0

	for n != 0 {
		n_prev = n
		fmt.Scan(&n)
		if n > n_prev {
			count++
		}
	}

	fmt.Println(count)
}
