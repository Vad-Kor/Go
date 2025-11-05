package main

import (
	"fmt"
)

func main_351() {
	var sex, result string
	var age int

	fmt.Scan(&age, &sex)

	ok := (12 <= age) && (age <= 18) && (sex == "m")

	if ok {
		result = "YES"
	} else {
		result = "NO"
	}

	fmt.Println(result)
}
