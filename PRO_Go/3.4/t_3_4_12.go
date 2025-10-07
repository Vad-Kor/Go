package main

import (
	"fmt"
)

func main_3412() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	exist := (a+b > c) && (a+c > b) && (b+c > a)

	if exist {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
