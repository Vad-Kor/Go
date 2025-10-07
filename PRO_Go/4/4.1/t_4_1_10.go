package main

import (
	"fmt"
)

func main_4110() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		fmt.Println(i)
	}
}
