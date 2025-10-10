package main

import (
	"fmt"
)

func main_435() {
	var n, f int

	fmt.Scan(&n)
	f = 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			f *= i
		}
	}
	fmt.Println(f)
}
