package main

import (
	"fmt"
	"strconv"
)

func main() {
	// welcome message
	fmt.Println("**Welcome to the calculator program!**\n")

	// num1 input and validation
	var num1 string
	fmt.Print("Please enter your first number: ")
	fmt.Scanln(&num1)

	fnum1, err := strconv.ParseFloat(num1, 64)
	for err != nil {
		fmt.Println("Invalid number")
		fmt.Print("Please enter your first number: ")
		fmt.Scanln(&num1)
		fnum1, err = strconv.ParseFloat(num1, 64)
	}

	// operator input and validation
	var validOperators = []string{"+", "-", "*", "/"}
	var operator string
	fmt.Print("Please enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	for contains(validOperators, operator) == false {
		fmt.Println("Invalid operator")
		fmt.Print("Please enter an operator (+, -, *, /): ")
		fmt.Scanln(&operator)
	}

	// num2 input and validation
	var num2 string
	fmt.Print("Please enter your second number: ")
	fmt.Scanln(&num2)

	fnum2, err := strconv.ParseFloat(num2, 64)
	for err != nil {
		fmt.Println("Invalid number")
		fmt.Print("Please enter your second number: ")
		fmt.Scanln(&num2)
		fnum2, err = strconv.ParseFloat(num2, 64)
	}
	if operator == "/" && fnum2 == 0 {
		fmt.Println("Cannot divide by zero")
		return
	} else {
		fmt.Printf("%v %v %v = %v\n", num1, operator, num2, calculate(fnum1, fnum2, operator))
	}
}

// contains checks if a given item exists in a slice of strings.
// It returns true if the item is found, otherwise it returns false.
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// calculate performs arithmetic operations on two numbers based on the given operator.
// It returns the result of the operation.
func calculate(num1, num2 float64, operator string) float64 {
	switch operator {
	case "+":
		return add(num1, num2)
	case "-":
		return subtract(num1, num2)
	case "*":
		return multiply(num1, num2)
	case "/":
		return divide(num1, num2)
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	result := num1 / num2
	return result
}
