package main

import (
	"fmt"
)

func main_4159() {
	var y int
	fmt.Scan(&y)
	result := false
	nextY := y

	for !result {
		nextY++
		d1 := nextY / 1000
		d2 := (nextY % 1000) / 100
		d3 := (nextY / 10) % 10
		d4 := nextY % 10
		if d1 != d2 && d1 != d3 && d1 != d4 && d2 != d3 && d2 != d4 && d3 != d4 {
			result = true
		}
	}

	fmt.Println(result)
}
