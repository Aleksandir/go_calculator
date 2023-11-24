package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
	// take the string and parse it into a slice of strings, separated by valid operators

	// order of operations: brackets, exponents, multiplication, division, addition, subtraction (BEDMAS)
	// loop through the slice and perform operations in order of operations
	// if the operator is a bracket, call equation_parser on the string inside the brackets
	// if the operator is an exponent, call square on the number before the operator and the number after the operator
	// remove the number before the operator and the operator from the slice
	// replace the number after the operator with the result of the square function
	// if the operator is multiplication, call multiply on the number before the operator and the number after the operator
	// if the operator is division, call divide on the number before the operator and the number after the operator
	// if the operator is addition, call add on the number before the operator and the number after the operator
	// if the operator is subtraction, call subtract on the number before the operator and the number after the operator

	for i, a := range equation {
		if string(a) == "(" {
			closingBracketIndex := strings.Index(equation[i:], ")")
			if closingBracketIndex == -1 {
				fmt.Println("Error: mismatched brackets")
				break
			}
			substring := equation[i+1 : i+closingBracketIndex]
			equation_parser(substring)
		}

		if string(i) == "^" || string(i) == "**" {
			// call square on the number before the operator and the number after the operator

		}
		if string(i) == "*" {
			// call multiply on the number before the operator and the number after the operator
		}
		if string(i) == "/" {
			// call divide on the number before the operator and the number after the operator
		}
		if string(i) == "+" {
			// call add on the number before the operator and the number after the operator
		}
		if string(i) == "-" {
			// call subtract on the number before the operator and the number after the operator
		}
	}

	// will return a slice of strings
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
