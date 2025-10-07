package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	separator := scanner.Text()
	_ = scanner.Scan()
	str1 := scanner.Text()
	_ = scanner.Scan()
	str2 := scanner.Text()
	_ = scanner.Scan()
	str3 := scanner.Text()

	fmt.Print(str1 + separator + str2 + separator + str3 + separator)
}
