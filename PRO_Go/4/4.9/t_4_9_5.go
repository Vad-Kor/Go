package main

import (
	"fmt"
)

func main_495() {
	count := 0

	for i := 100; i <= 999; i++ {
		if i%10+i/100+(i/10)%10 == 8 {
			count++
		}
	}
	fmt.Println(count)
}
