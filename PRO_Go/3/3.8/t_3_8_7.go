package main

import (
	"fmt"
)

func main_387() {
	var rentalCost, freeMinutes, actualTripDuration, additionalMinuteCost int

	fmt.Scan(&rentalCost, &freeMinutes, &actualTripDuration, &additionalMinuteCost)

	additionalCost := (actualTripDuration - freeMinutes) * additionalMinuteCost

	if additionalCost < 0 {
		additionalCost = 0
	}

	totalCost := rentalCost + additionalCost

	fmt.Println(totalCost)
}
