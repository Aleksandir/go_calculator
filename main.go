package main

import (
	"fmt"
	"math"
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
	var validOperators = []string{"+", "-", "*", "/", "^", "**", "%"}
	var operator string
	fmt.Printf("Please enter an operator %v: ", validOperators)
	fmt.Scanln(&operator)
	for contains(validOperators, operator) == false {
		fmt.Println("Invalid operator")
		fmt.Printf("Please enter an operator (%s): ", validOperators)
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
	// may not need this after new module addition
	if operator == "/" && fnum2 == 0 {
		fmt.Println("Cannot divide by zero")
		return
	} else {
		expression := fmt.Sprintf("%v %v %v", fnum1, operator, fnum2)
		result := calculate(fnum1, fnum2, operator)
		fmt.Printf("%s = %v\n", expression, result)
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
	case "^", "**":
		return square2(num1)
	case "%":
		return remainder(num1, num2)
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

// func equation_parser(equation string) []string {
// 	var parsedEquation []string
// 	tempNumber := ""

// 	for _, char := range equation {
// 		if unicode.IsDigit(char) {
// 			tempNumber += string(char)
// 		} else {
// 			if tempNumber != "" {
// 				parsedEquation = append(parsedEquation, tempNumber)
// 				tempNumber = ""
// 			}
// 			parsedEquation = append(parsedEquation, string(char))
// 		}
// 	}

// 	if tempNumber != "" {
// 		parsedEquation = append(parsedEquation, tempNumber)
// 	}

// 	fmt.Println(parsedEquation)
// 	return parsedEquation
// }
// func perform_operations(parsedEquation []string) int {
// 	// This is a placeholder function. You'll need to implement this function to perform the actual operations.
// 	// For now, it just converts the first element of parsedEquation to an integer and returns it.
// 	if len(parsedEquation) > 0 {
// 		result, _ := strconv.Atoi(parsedEquation[0])
// 		return result
// 	}
// 	return 0
// }

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
func square2(num1 float64) float64 {
	return num1 * num1
}
func square(num1, pow float64) float64 {
	return math.Pow(num1, pow)
}

func remainder(num1, num2 float64) float64 {
	return math.Mod(num1, num2)
}
