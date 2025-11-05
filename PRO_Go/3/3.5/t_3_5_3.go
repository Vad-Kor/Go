package main

import (
	"fmt"
)

func main_353() {
	var result string
	var n int
	var ok bool = true

	fmt.Scan(&n)

	if (n%2 == 0) && (20 < n || 2 <= n && n <= 5) {
		ok = false
	}

	if ok {
		result = "YES"
	} else {
		result = "NO"
	}

	fmt.Println(result)
}
