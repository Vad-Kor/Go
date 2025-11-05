package main

import (
	"fmt"
)

func main() {
	var x, y int
	var result string

	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		result = "1"
	} else if x < 0 && y > 0 {
		result = "2"
	} else if x < 0 && y < 0 {
		result = "3"
	} else {
		result = "4"
	}

	fmt.Println(result)
}
