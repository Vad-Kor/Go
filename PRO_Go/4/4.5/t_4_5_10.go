package main

import (
	"fmt"
	"strconv"
)

func main_4510() {
	var n int
	var result string

	fmt.Scan(&n)

	for n > 0 {
		result += strconv.Itoa(n % 10)
		n = n / 10
	}

	fmt.Println(result)
}
