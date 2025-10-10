package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	i := 2
	for n%i != 0 {
		i++

	}
	fmt.Println(n / i)
}
