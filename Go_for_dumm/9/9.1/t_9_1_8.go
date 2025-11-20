package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main_918() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	input := myscanner.Text()

	stringNumbers := strings.Split(input, " ")
	var numbers []int
	for _, stringNumber := range stringNumbers {
		number, _ := strconv.Atoi(stringNumber)
		numbers = append(numbers, number)
	}

	var maxValue, prevMaxValue, maxValueIndex, prevMaxValueIndex int

	if numbers[0] >= numbers[1] {
		maxValue = numbers[0]
		maxValueIndex = 0
		prevMaxValue = numbers[1]
		prevMaxValueIndex = 1
	} else {
		maxValue = numbers[1]
		maxValueIndex = 1
		prevMaxValue = numbers[0]
		prevMaxValueIndex = 0
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > maxValue {
			prevMaxValue = maxValue
			prevMaxValueIndex = maxValueIndex
			maxValue = numbers[i]
			maxValueIndex = i
		} else if numbers[i] > prevMaxValue {
			prevMaxValue = numbers[i]
			prevMaxValueIndex = i
		}
	}

	if maxValueIndex <= prevMaxValueIndex {
		fmt.Println(maxValue, prevMaxValue)
	} else {
		fmt.Println(prevMaxValue, maxValue)
	}
}
