package main

import (
	"fmt"
)

func main_4152() {
	var n, nI int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&nI)
		if nI == 0 {
			count++
		}

	}
	fmt.Println(count)
}
