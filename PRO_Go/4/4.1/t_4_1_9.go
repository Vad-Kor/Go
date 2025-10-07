package main

import (
	"fmt"
)

func main_419() {
	var x int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}
