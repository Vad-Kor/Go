package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()

	X, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Следующее за числом", X, "число:", X+1)
	fmt.Println("Для числа", X, "предыдущее число:", X-1)
	fmt.Println("row1")
}
