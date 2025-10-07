package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1() {
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	password := scanner.Text()
	_ = scanner.Scan()
	password_repeat := scanner.Text()

	var result string

	if password == password_repeat {
		result = "Пароль принят"
	} else {
		result = "Пароль не принят"
	}
	fmt.Println(result)
}
