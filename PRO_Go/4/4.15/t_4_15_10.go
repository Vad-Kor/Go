package main

import (
	"fmt"
)

func main_41510() {
	var n int
	fmt.Scan(&n)

	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			fmt.Printf("%d ", col*row)
		}
		fmt.Println()
	}
	if n == 0 {
		fmt.Println("Таблица пустая")
	}
}
