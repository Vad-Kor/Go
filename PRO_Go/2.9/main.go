package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var n1, n2, n3 float64

	fmt.Scan(&n)
	n = int(math.Abs(float64(n)))
	n1 = float64(n % 10)
	n2 = float64((n / 10) % 10)
	n3 = float64(n / 100)

	fmt.Println(math.Max(n1, math.Max(n2, n3)))

}
