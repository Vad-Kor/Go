package main

import (
	"fmt"
)

func main_4411() {
	var n int

	fmt.Scan(&n)
	step := 1

	for step <= n {
		fmt.Println(step)
		step = step * 2
	}
}
