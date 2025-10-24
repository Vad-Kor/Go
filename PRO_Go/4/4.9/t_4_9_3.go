package main

import (
	"fmt"
)

func main_493() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	first := true
	for x := 0; x <= 1000; x++ {

		if a*x*x*x+b*x*x+c*x+d == 0 {
			if first {
				first = false
			} else {
				fmt.Print(" ")
			}
			fmt.Print(x)
		}
	}
}
