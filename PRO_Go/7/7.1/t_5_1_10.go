package main

import (
	"fmt"
)

func main_41510() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(numbers)
	for i := 1; i < size; i += 2 {
		tmp := numbers[i]
		numbers[i] = numbers[i-1]
		numbers[i-1] = tmp
	}
	fmt.Println(numbers)
}
