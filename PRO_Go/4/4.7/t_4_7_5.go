package main

import (
	"fmt"
)

func main_475() {
	var n int
	var simpleNumber int = 2

	fmt.Scan(&n)
	count := 0

	for n != 1 {

		if n%simpleNumber == 0 {
			if count != 0 {
				fmt.Print(" ")
			}
			fmt.Print(simpleNumber)
			count++
			n = n / simpleNumber
		} else {
			simpleNumber++
		}
	}
}
