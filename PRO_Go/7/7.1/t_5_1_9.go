package main

import (
	"fmt"
)

func main_4159() {
	var size, d int
	fmt.Scan(&size)
	var numbers []int

	count := 0

	for i := 0; i < size; i++ {
		fmt.Scan(&d)
		numbers = append(numbers, d)
		if i != 0 && numbers[i] > numbers[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
