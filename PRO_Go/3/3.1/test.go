package main

import (
	"fmt"
)

func main() {
	var x int
	var result int

	fmt.Scan(&x)

	if x < 0 {
		result = -1
	} else if x == 0 {
		result = 0
	} else {
		result = 1
	}

	fmt.Println(result)
}
