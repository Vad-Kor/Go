package main

import (
	"fmt"
)

func get_class_count(count int) (result int) {
	result = count/2 + count%2

	return result
}

func main_3413() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(get_class_count(a) + get_class_count(b) + get_class_count(c))
}
