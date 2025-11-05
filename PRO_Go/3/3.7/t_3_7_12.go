package main

import (
	"fmt"
)

func main_3712() {
	var n int

	fmt.Scan(&n)
	n1 := n / 100
	n2 := (n % 100) / 10
	n3 := n % 10
	maxN := max(n1, n2, n3)
	minN := min(n1, n2, n3)
	midN := n1
	if minN < n1 && n1 < maxN {
		midN = n1
	}
	if minN < n2 && n2 < maxN {
		midN = n2
	}
	if minN < n3 && n3 < maxN {
		midN = n3
	}

	result := "Число неинтересное"

	if minN+maxN == midN {
		result = "Число интересное"
	}

	fmt.Println(result)
}
