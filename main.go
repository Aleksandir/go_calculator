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
	if operator == "/" && fnum2 == 0 {
		fmt.Println("Cannot divide by zero")
		return
	} else {
		equation_parser(num1 + operator + num2)
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
	case "^", "**":
		return square2(num1)
	case "%":
		return remainder(num1, num2)
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

// take a string and parse it into a slice of strings, then return the slice
// this is so the user can input a string of numbers and operators and the program will parse it
func equation_parser(equation string) []string {
	// 1. Define a function equation_parser that takes a string as input.
	// 2. Initialize an empty slice to hold the parsed elements of the equation.
	// 3. Loop over the characters in the input string.
	// 4. If the character is a digit, keep adding to a temporary string until a non-digit character is found. This handles multi-digit numbers.
	// 5. If the character is a non-digit (i.e., an operator), add the number accumulated so far to the slice, add the operator to the slice, and reset the temporary string.
	// 6. After the loop, add any remaining number to the slice.
	// 7. Now you have a slice with numbers and operators as separate elements.
	// 8. Loop over the slice and perform the operations in the correct order (following BEDMAS/BODMAS rules). You might need to implement separate functions for each operation.
	// 9. Return the final result.


	// below is sudo code
    parsed_equation = []
    temp_number = ""
    for char in equation:
        if char is a digit:
            temp_number += char
        else:
            if temp_number is not empty:
                parsed_equation.append(temp_number)
                temp_number = ""
            parsed_equation.append(char)
    if temp_number is not empty:
        parsed_equation.append(temp_number)
    result = perform_operations(parsed_equation)
    return result
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
func square2(num1 float64) float64 {
	return num1 * num1
}
func square(num1, pow float64) float64 {
	return math.Pow(num1, pow)
}

func remainder(num1, num2 float64) float64 {
	return math.Mod(num1, num2)
}
