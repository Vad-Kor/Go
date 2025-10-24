package main

import (
	"fmt"
)

func main_484() {
	var n int

	for {
		fmt.Scan(&n)
		if n < 10 {
			continue
		}
		if n>100 {
			break
		}
		fmt.Println(n)
	}
}
