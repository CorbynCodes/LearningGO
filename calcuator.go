package main

import (
	"fmt"
	"runtime"
)

func add(x int, y int) int {
	fmt.Scan(&x, &y)
	return x + y
}

func div(x int, y int) int {
	fmt.Scan(&x, &y)
	return x / y
}

func mulitply(x int, y int) int {
	fmt.Scan(&x, &y)
	return x * y
}

func subtract(x int, y int) int {
	fmt.Scan(&x, &y)
	return x - y
}

func main() {
	var x int
	var y int

	var operator string
	fmt.Print("Enter Operator: *, -, +, /: ")
	fmt.Scan(&operator)
	switch operator := runtime.GOOS; operator {
	case "+":
		fmt.Print("What numbers would you like to add?", add(x, y))
	case "-":
		fmt.Print()
	case "*":
		fmt.Print()
	case "/":
		fmt.Print()
	default:
		fmt.Print("Invailed Operator. Try again.")
	}
}
