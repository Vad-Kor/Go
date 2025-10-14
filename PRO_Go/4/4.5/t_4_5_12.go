package main

import (
	"fmt"
	"strconv"
)

func main_4512() {
	var n int
	var result string

	fmt.Scan(&n)

	for n > 0 {
		result += strconv.Itoa(n % 2)
		n = n / 2
	}

	fmt.Println(result)
}
