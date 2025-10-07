package main

import (
	"fmt"
)

func main_354() {
	var n float64

	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		fmt.Println(1.0 / n)
	}
}
