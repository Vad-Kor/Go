package main

import (
	"fmt"
)

func main_458() {
	var n int

	fmt.Scan(&n)
	count := 0

	for n > 0 {
		if n%10 == 4 {
			count++
		}
		n = n / 10

	}

	fmt.Println(count)
}
