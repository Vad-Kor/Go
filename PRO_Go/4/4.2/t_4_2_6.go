package main

import (
	"fmt"
)

func main_426() {
	var n, x int
	fmt.Scan(&n)

	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		sum += x
	}
	fmt.Println(sum)
}
