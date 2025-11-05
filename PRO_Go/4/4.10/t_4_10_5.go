package main

import (
	"fmt"
)

func main_4105() {
	var n, nI int
	fmt.Scan(&n)
	fmt.Scan(&nI)
	max := nI
	min := nI

	for i := 1; i < n; i++ {
		fmt.Scan(&nI)
		if nI < min {
			min = nI
		}
		if nI > max {
			max = nI
		}
	}
	fmt.Println(max - min)
}
