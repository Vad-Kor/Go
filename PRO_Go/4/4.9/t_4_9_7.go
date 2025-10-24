package main

import "fmt"

func main_497() {
	count := 0
	n := 1
	for count < 3 {
		sum := 0
		for i := 1; i < n; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		if sum == n {
			fmt.Println(n, ",")
			count++
		}
		n++
	}

}
