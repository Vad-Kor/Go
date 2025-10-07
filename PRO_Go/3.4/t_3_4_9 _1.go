package main

import (
	"fmt"
)

func main_3491() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)

	n256 := min(k2, k5, k6)
	n32 := min(k2-n256, k3)
	sum := 256*n256 + n32*32

	fmt.Println(sum)
}
