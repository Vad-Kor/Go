package main

import (
	"fmt"
)

func main_4157() {
	var n int
	fmt.Scan(&n)

	sum := 0

	for sum > 9 || sum == 0 {
		sum = 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}

	fmt.Println(sum)
}
