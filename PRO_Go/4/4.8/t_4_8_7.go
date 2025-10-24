package main

import (
	"fmt"
	"strings"
)

func main_487() {
	var s, s1 string

	fmt.Scan(&s)

	i := 0
	for {
		fmt.Scan(&s1)
		i++
		if len(s1) >= len(s) && strings.Contains(s1, s) {
			fmt.Println(i)
			break
		}
	}
}
