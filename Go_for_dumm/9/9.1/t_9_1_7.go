package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main_917() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	input := myscanner.Text()

	stringWords := strings.Split(input, " ")
	for i, j := 0, len(stringWords)-1; i < j; i, j = i+1, j-1 {
		stringWords[i], stringWords[j] = stringWords[j], stringWords[i]
	}

	result := strings.Trim(fmt.Sprint(stringWords), "[]")
	fmt.Println(result)
}
