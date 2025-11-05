package main

import (
	"fmt"
)

func main_348() {
	var k, m int
	fmt.Scan(&k, &m)
	d := k / m
	if k%m != 0 {
		d++
	}
	fmt.Println(d)
}
