package main

import (
	"fmt"
)

func main_375() {
	var n int

	fmt.Scan(&n)

	result := n

	if n%2 == 0 {
		result += 2
	} else {
		result += 1
	}

	fmt.Println(result)
}
