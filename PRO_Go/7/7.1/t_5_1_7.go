package main

import (
	"fmt"
	"strings"
)

func main_4157() {
	var size, d int
	fmt.Scan(&size)
	var numbers []int

	for i := 0; i < size; i++ {
		fmt.Scan(&d)
		if i%3 == 0 {
			numbers = append(numbers, d)
		}
	}
	result := strings.Trim(fmt.Sprint(numbers), "[]")
	fmt.Println(result)
}
