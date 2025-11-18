package main

import (
	"fmt"
)

func main_4155() {
	var a, n int
	fmt.Scan(&a)
	fmt.Scan(&n)
	res := 1

	for n > 0 {
		res *= a
		n--
	}

	fmt.Println(res)
}
