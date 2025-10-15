package main

import (
	"fmt"
)

func main_474() {
	var n int

	fmt.Scan(&n)

	maxSum := 0
	minN := 0

	for i := 1; i <= n; i++ {
		j := i
		curSum := 0
		for j != 0 {
			curSum += j % 10
			j /= 10
		}
		if curSum > maxSum {
			maxSum = curSum
			minN = i
		}

	}

	fmt.Println(minN, maxSum)
}
