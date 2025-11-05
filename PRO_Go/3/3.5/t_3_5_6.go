package main

import (
	"fmt"
)

func main_356() {
	var d1, d2, d3 int

	fmt.Scan(&d1, &d2, &d3)

	result := min(2*d1+2*d2, d1+d2+d3, 2*d1+2*d3, 2*d2+2*d3)

	fmt.Println(result)
}
