package main

import (
	"fmt"
)

func main_4412() {
	var a, b, f int

	fmt.Scan(&a, &b)
	f = 1

	for i := a; i <= b; i++ {
		if i%10 == 7 || i%10 == -7 {
			f *= i
		}
	}
	fmt.Println(f)
}
