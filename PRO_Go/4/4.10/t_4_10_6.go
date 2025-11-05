package main

import (
	"fmt"
)

func main_4106() {
	var n, nI int
	fmt.Scan(&n)
	fmt.Scan(&nI)
	max := nI
	min := nI
	result := "NO"

	for i := 1; i < n; i++ {
		fmt.Scan(&nI)
		if nI > max {
			max = nI
		}
		if nI < min {
			min = nI
		}
	}
	if min < 30 {
		result = "YES"
	}
	fmt.Println(max, result)
}
