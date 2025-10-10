package main

import (
	"fmt"
)

func main_428() {
	var n, x int
	var result string = "NO"
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 0 && result == "NO" {
			result = "YES"
		}
	}
	fmt.Println(result)
}
