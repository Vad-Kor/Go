package main

import (
	"fmt"
	"math"
)

func main_3411() {
	var qx, qy, ox, oy int
	var result string
	fmt.Scan(&qx, &qy, &ox, &oy)

	pol := (qx == ox) || (qy == oy) || (math.Abs(float64(qx-ox)) == math.Abs(float64(qy-oy)))

	if pol {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println(result)
}
