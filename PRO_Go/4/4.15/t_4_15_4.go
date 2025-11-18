package main

import (
	"fmt"
	"math"
)

func main_4154() {
	var n int
	fmt.Scan(&n)
	res := 0
	d := 0
	for n > 0 {
		dig := n % 10
		if dig != 5 && dig != 7 {
			res = res + dig*int(math.Pow10(d))
			d++
		}
		n /= 10
	}

	fmt.Println(res)
}
