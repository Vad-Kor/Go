package main

import (
	"fmt"
)

func main_459() {
	var n int

	fmt.Scan(&n)
	sum := 0
	n_o := n

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	if n_o%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
