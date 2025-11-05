package main

import (
	"fmt"
)

func main_374() {
	var n int

	fmt.Scan(&n)

	result := "NO"

	if (n%10)%2 == 0 {
		result = "YES"
	}

	fmt.Println(result)
}
