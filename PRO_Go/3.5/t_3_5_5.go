package main

import (
	"fmt"
)

func main_355() {
	var result string
	var n int

	fmt.Scan(&n)

	switch {
	case n <= 3:
		result = "начинающий"
	case n <= 7:
		result = "младший разработчик"
	case n <= 15:
		result = "средний разработчик"
	default:
		result = "старший разработчик"
	}

	fmt.Println(result)
}
