package main

import "fmt"

func main_c() {
	var a, b float32
	var op, result string

	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		result = fmt.Sprintf("%g", a+b)
	case "-":
		result = fmt.Sprintf("%g", a-b)
	case "*":
		result = fmt.Sprintf("%g", a*b)
	case "/":
		if b == 0 {
			result = "Делить на ноль нельзя!"
		} else {
			result = fmt.Sprintf("%g", a/b)
		}
	case "%":
		if b == 0 {
			result = "На ноль делить нельзя!!"
		} else {
			result = fmt.Sprintf("%g", int(a)%int(b))
		}
	default:
		result = "Неверная операция"
	}

	fmt.Println(result)
}
