package main

import (
	"fmt"
)

func main_4102() {
	var n, nI int
	fmt.Scan(&n)
	fmt.Scan(&nI)
	min := nI

	for i := 1; i < n; i++ {
		fmt.Scan(&nI)
		if nI < min {
			min = nI
		}
	}
	fmt.Println(min)
}
