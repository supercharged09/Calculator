package main

import "fmt"

func main() {
	var a, b float64
	var operator string
	fmt.Println("Введите число 1:")
	fmt.Scan(&a)
	fmt.Println("Введите оператор (+,-,*./): ")
	fmt.Scan(&operator)
	fmt.Println("Введите число 2:")
	fmt.Scan(&b)
	fmt.Print(a, operator, b, "=")
	switch operator {
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "*":
		fmt.Print(a * b)
	case "/":
		fmt.Print(a / b)
	default:
		fmt.Print("invalid operator/number")
	}
}
