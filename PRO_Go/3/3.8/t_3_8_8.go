package main

import (
	"fmt"
)

func main_388() {
	var ticketPrice, windowSeatCost, inFlightMealsCost, carryOnBaggageWeight int

	fmt.Scan(&ticketPrice, &windowSeatCost, &inFlightMealsCost, &carryOnBaggageWeight)

	extraCost := 0
	if carryOnBaggageWeight > 3 {
		extraCost = 200 * (carryOnBaggageWeight - 3)
	}

	totalCost := ticketPrice + windowSeatCost + inFlightMealsCost + extraCost

	fmt.Printf("Полёт обойдётся в %d рублей", totalCost)
}
