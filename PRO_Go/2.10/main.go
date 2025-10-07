package main

import (
	"fmt"
)

func main() {
	var a, d float64

	fmt.Scan(&a, &d)

	fmt.Println(a - a*(d/100.0))
}
