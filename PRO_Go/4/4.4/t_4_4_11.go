package main

import (
	"fmt"
)

func main_4411() {
	var a, b, f int

	fmt.Scan(&a, &b)
	f = 1

	for i := a; i <= b; i++ {
		f *= i
	}
	fmt.Println(f)
}
