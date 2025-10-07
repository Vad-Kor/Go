package main

import (
	"fmt"
	"math"
)

func main_3413() {
	var n int
	fmt.Scan(&n)
	pol := true
	for i := 1; i <= 2; i++ {
		a := (n / int(math.Pow(10, float64(i-1)))) % 10
		b := (n / int(math.Pow(10, float64(4-i)))) % 10
		if a == b {
			continue
		} else {
			pol = false
			break
		}
	}
	if pol {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
