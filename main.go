package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Get the equation from the query parameters
	equation := r.URL.Query().Get("equation")

	// Parse and evaluate the equation
	expression, err := govaluate.NewEvaluableExpression(equation)
	if err != nil {
		http.Error(w, "Invalid equation", http.StatusBadRequest)
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		http.Error(w, "Error evaluating equation", http.StatusInternalServerError)
		return
	}

	// Send the result back to the client
	fmt.Fprintf(w, "%v = %v\n", equation, result)
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "calculator.html")
	})
	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	// welcome message
// 	fmt.Println("**Welcome to the calculator program!**\n")
// 	for {
// 		// user input
// 		var equation string
// 		fmt.Print("Please enter your equation: ")
// 		fmt.Scanln(&equation)
// 		if equation == "q" {
// 			fmt.Println("Quitting...")
// 			return
// 		}

// 		// parse equation
// 		expression, err := govaluate.NewEvaluableExpression(equation)
// 		if err != nil {
// 			fmt.Println("Error parsing equation")
// 			fmt.Println(err)
// 			return
// 		}

// 		// evaluate equation
// 		result, err := expression.Evaluate(nil)
// 		if err != nil {
// 			fmt.Println("Error evaluating equation")
// 			fmt.Println(err)
// 			return
// 		}

// 		// round the result
// 		floatResult, ok := result.(float64)
// 		if ok {
// 			floatResult := floatResult * 100
// 			result = math.Round(floatResult) / 100
// 		}

// 		// print result
// 		fmt.Printf("%v = %v\n", equation, result)
// 	}
// }
