package main

import "fmt"

func main_838() {
	var n int
	for {
		fmt.Scan(&n)
		if n != 0 && n != 1 {
			d := false
			for i := 2; i < n; i++ {
				if n%i == 0 {
					d = true
				}
			}
			if d == false {
				break
			}
		}
	}
	fmt.Println(n)
}
