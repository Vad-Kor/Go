package main

import (
	"fmt"
	"math"
)

func main_337() {
	// var result string = "Нет результата"
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	a = 2 * a
	b = -b
	d := math.Pow(b, 2) - 2*a*c
	if d != 0 {
		d = math.Sqrt(d)
	}
	if d == 0 {
		fmt.Println(b / a)
	}
	if d > 0 {
		if a > 0 {
			fmt.Println((b - d) / a)
			fmt.Println((b + d) / a)
		} else {
			fmt.Println((b + d) / a)
			fmt.Println((b - d) / a)
		}
	}

}
