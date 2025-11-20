package main

import (
	"fmt"
	"strings"
)

func main_415102() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
		if i%2 != 0 {
			tmp := numbers[i]
			numbers[i] = numbers[i-1]
			numbers[i-1] = tmp
		}
	}

	result := strings.Trim(fmt.Sprint(numbers), "[]")
	fmt.Println(result)
}
