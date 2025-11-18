package main

import (
	"fmt"
)

func main_4158() {
	var a, b int
	fmt.Scan(&a, &b)
	result := "NO"

	for i := b; i >= a; i-- {
		ii := i
		if ii < 0 {
			ii *= -1
		}
		if ii%7 == 0 {
			result = fmt.Sprintf("%d", i)
			break
		}
	}

	fmt.Println(result)
}
