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

/* package main

import "fmt"

func main() {
	var a, b float64
	var operator string
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите оператор (+,-,*,/")
	fmt.Scan(&operator)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	fmt.Print(a, operator, b, " = ")

	if operator == "+" {
		fmt.Print(a + b)
	} else if operator == "-" {
		fmt.Print(a - b)
	} else if operator == "*" {
		fmt.Print(a * b)
	} else if operator == "/" {
		fmt.Print(a / b)
	} else {
		fmt.Print("invalid operator/number")
	}
} */
