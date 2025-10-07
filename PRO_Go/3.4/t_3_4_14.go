package main

import (
	"fmt"
)

func main_3414() {
	var w int
	fmt.Scan(&w)

	result := w%2 == 0
	fmt.Println(w, w/2)

	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
