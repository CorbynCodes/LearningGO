package main

import (
	"fmt"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

func div(x int, y int) int {
	return x / y
}

func mulitply(x int, y int) int {
	return x * y
}

func subtract(x int, y int) int {
	return x - y
}

func main() {
	var x, y int

	var operator string
	fmt.Print("Enter Operator: *, -, +, /: ")
	fmt.Scan(&operator)
	switch operator := runtime.GOOS; operator {
	case "+":
		fmt.Print("What numbers would you like to add?: ")
		fmt.Scan(&x, &y)
		fmt.Print("Result: ", add(x, y))
	case "-":
		fmt.Print("What numbers would you like to subtract: ")
		fmt.Scan(&x, &y)
		fmt.Print("Result: ", subtract(x, y))
	case "*":
		fmt.Print("What numbers would you like to muiltply: ")
		fmt.Scan(&x, &y)
		fmt.Print("Result: ", mulitply(x, y))
	case "/":
		fmt.Print("What numbers would you like to divide: ")
		fmt.Scan(&x, &y)
		fmt.Print("Result,", div(x, y))
	default:
		fmt.Print("Invailed Operator. Try again.")
	}
}
