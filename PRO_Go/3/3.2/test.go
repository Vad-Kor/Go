package main

import (
	"fmt"
)

func main1() {
	var n int
	var res int = 0

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		res += i
	}

	fmt.Println(res)

}
