package main

import "fmt"

func main_837() {
	var n, f int
	f_1 := 0
	f_2 := 1

	fmt.Scan(&n)

	for {
		f = f_1 + f_2
		if f > n {
			break
		}
		f_1 = f_2
		f_2 = f
	}
	fmt.Println(f)
}
