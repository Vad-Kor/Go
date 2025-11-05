package main

import (
	"fmt"
)

func main_3310() {
	const (
		v1 = 28
		v2 = 30
		v3 = 31
	)
	var a int
	fmt.Scan(&a)

	switch a {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(v3)
	case 2:
		fmt.Println(v1)
	case 4, 6, 9, 11:
		fmt.Println(v2)
	}
}
