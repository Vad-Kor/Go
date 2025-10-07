package main

import "fmt"

func main_835() {
	var n, a int
	fmt.Scan(&n)
	sum := 0
	for {
		fmt.Scan(&a)
		sum += a
		if sum > n {
			break
		}
	}
	fmt.Println(sum)
}
