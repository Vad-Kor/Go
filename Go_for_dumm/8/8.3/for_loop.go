package main

import "fmt"

func main_834() {
	var n int
	sum := 0
	for {
		fmt.Scan(&n)
		if n < 0 {
			break
		}
		sum += n
	}
	fmt.Println(sum)
}
