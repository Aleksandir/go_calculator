package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	// welcome message
	fmt.Println("**Welcome to the calculator program!**\n")
	for {
		// user input
		var equation string
		fmt.Print("Please enter your equation: ")
		fmt.Scanln(&equation)
		if equation == "q" {
			fmt.Println("Quitting...")
			return
		}

		// parse equation
		expression, err := govaluate.NewEvaluableExpression(equation)
		if err != nil {
			fmt.Println("Error parsing equation")
			fmt.Println(err)
			return
		}

		// evaluate equation
		result, err := expression.Evaluate(nil)
		if err != nil {
			fmt.Println("Error evaluating equation")
			fmt.Println(err)
			return
		}
		fmt.Printf("%v = %v\n", equation, result)
	}
}
