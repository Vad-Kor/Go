package main

import "fmt"

func main_836() {
	var n int
	count := 0
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n%2 == 0 {
			count++
		} else {
			fmt.Println("Нечётное число!")
		}
	}
	fmt.Println(count)
}
