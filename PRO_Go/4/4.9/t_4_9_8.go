package main

import (
	"fmt"
)

func main_498() {
	var a, b, c, d, e int

	fmt.Scan(&a, &b, &c, &d, &e)

	count := 0
	for x := 0; x <= 1000; x++ {

		if (x != e) && ((a*x*x*x + b*x*x + c*x + d) == 0) {
			count++
		}
	}
	fmt.Println(count)
}
