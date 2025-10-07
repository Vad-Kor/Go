package main

import (
	"fmt"
)

func main_349() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)
	sum := 0
	for {
		if k2 > 0 && k5 > 0 && k6 > 0 {
			sum += 256
			k2--
			k5--
			k6--
		} else if k2 > 0 && k3 > 0 {
			sum += 32
			k2--
			k3--
		} else {
			break
		}
	}
	fmt.Println(sum)
}
