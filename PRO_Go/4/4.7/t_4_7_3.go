package main

import (
	"fmt"
)

func main_473() {
	var n int

	fmt.Scan(&n)
	count := 0

	for i := 1; i <= n; i++ {
		j := i
		for j != 0 {
			if j%10 == 7 {
				count++
			}
			j /= 10
		}
	}

	fmt.Println(count)
}
