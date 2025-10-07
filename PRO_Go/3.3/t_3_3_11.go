package main

import (
	"fmt"
)

func main() {
	const (
		s1 = "Зима"
		s2 = "Весна"
		s3 = "Лето"
		s4 = "Осень"
	)
	var result string
	var a int
	fmt.Scan(&a)

	switch a {
	case 12, 1, 2:
		result = s1
	case 3, 4, 5:
		result = s2
	case 6, 7, 8:
		result = s3
	default:
		result = s4
	}
	fmt.Println(result)
}
