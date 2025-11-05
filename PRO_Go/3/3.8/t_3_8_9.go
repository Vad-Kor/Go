package main

import (
	"fmt"
)

func main_389() {
	var n int
	var end string

	fmt.Scan(&n)

	if (n%10 == 1) && (n%100 != 11) {
		end = ""
	} else if ((n%10 == 2) || (n%10 == 3) || (n%10 == 4)) && ((n%100 != 12) || (n%100 != 13) || (n%100 != 14)) {
		end = "а"
	} else {
		end = "ов"
	}

	fmt.Printf("%d Гимназист"+end, n)
}
