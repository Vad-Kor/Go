package main

import (
	"fmt"
)

func main_4103() {
	var n, nI int
	fmt.Scan(&n)
	fmt.Scan(&nI)
	max := nI

	for i := 1; i < n; i++ {
		fmt.Scan(&nI)
		if nI > max {
			max = nI
		}
	}
	fmt.Println(max)
}
