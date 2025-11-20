package main

import (
	"fmt"
)

func main_41512() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	result := "NO"

	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
		if i != 0 {
			if numbers[i] > 0 && numbers[i-1] > 0 || numbers[i] < 0 && numbers[i-1] < 0 {
				result = "YES"
			}
		}
	}
	fmt.Println(result)
}
