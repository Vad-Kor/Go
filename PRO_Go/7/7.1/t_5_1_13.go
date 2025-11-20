package main

import (
	"fmt"
	"strings"
)

func main_41513() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)

	for i := 1; i < size; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Scan(&numbers[0])
	result := strings.Trim(fmt.Sprint(numbers), "[]")
	fmt.Println(result)
}
