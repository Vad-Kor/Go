package main

import (
	"fmt"
)

func main_4153() {
	var n int
	fmt.Scan(&n)
	pol := 0
	m := n
	for m > 0 {
		pol = pol*10 + m%10
		m /= 10
	}

	if pol == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
