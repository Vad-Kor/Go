package main

import (
	"fmt"
)

func main_4413() {
	var a, b int

	fmt.Scan(&a, &b)

	nod := min(a, b)

	for a%nod != 0 || b%nod != 0 {
		nod--
	}
	fmt.Println(nod)
}
