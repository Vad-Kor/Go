package main

import (
	"fmt"
)

func main_352() {
	var result string
	var x int

	fmt.Scan(&x)

	ok := ((-30 < x) && (x <= -2)) || ((7 < x) && (x <= 25))

	if ok {
		result = "Принадлежит"
	} else {
		result = "Не принадлежит"
	}

	fmt.Println(result)
}
