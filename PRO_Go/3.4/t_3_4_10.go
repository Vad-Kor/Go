package main

import (
	"fmt"
)

func main_3410() {
	var hx, hy, ox, oy int
	var result string
	fmt.Scan(&hx, &hy, &ox, &oy)

	pol1 := (hy+2 == oy) && ((hx+1) == ox || (hx-1) == ox)
	pol2 := (hy+1 == oy) && ((hx+2) == ox || (hx-2) == ox)
	pol3 := (hy-2 == oy) && ((hx+1) == ox || (hx-1) == ox)
	pol4 := (hy-1 == oy) && ((hx+2) == ox || (hx-2) == ox)

	if pol1 || pol2 || pol3 || pol4 {
		result = "YES"
	} else {
		result = "NO"
	}
	fmt.Println(result)
}
