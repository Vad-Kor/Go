package main

import (
	"fmt"
)

func main_449() {
	var n int

	fmt.Scan(&n)
	count := 0

	for n%3 == 0 {
		n = n / 3
		count++
	}

	fmt.Println(count)
}
