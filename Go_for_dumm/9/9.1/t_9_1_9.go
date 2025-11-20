package main

import (
	"fmt"
)

func main_919() {
	const size = 15
	var numbers [size]int
	var evens []int
	var odds []int

	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	i, j := 0, 0 // Индексы для срезов evens и odds
	for i < len(evens) && j < len(odds) {
		fmt.Print(evens[i], " ")
		fmt.Print(odds[j], " ")
		i++
		j++
	}

	for ; i < len(evens); i++ {
		fmt.Print(evens[i], " ")
	}
	for ; j < len(odds); j++ {
		fmt.Print(odds[j], " ")
	}
}
