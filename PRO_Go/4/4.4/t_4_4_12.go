package main

import (
	"fmt"
)

func main_4412() {
	var n int

	fmt.Scan(&n)
	k := 0
	res := 1

	for res < n {
		k++
		res = res * 2

	}
	fmt.Println(k)
}
