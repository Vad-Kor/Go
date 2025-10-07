package main

import (
	"fmt"
)

func main() {
	var result string
	var a int
	fmt.Scan(&a)
	if (a%10)%2 == 0 {
		result = "YES"
	} else {
		result = "NO"
	}

	fmt.Println(result)
}
