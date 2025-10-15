package main

import (
	"fmt"
)

func main_465() {
	var n, prevN int

	fmt.Scan(&n)

	count := 0

	for n != 0 {
		prevN = n
		fmt.Scan(&n)
		if n*prevN < 0 {
			count++
		}
	}

	fmt.Println(count)
}
