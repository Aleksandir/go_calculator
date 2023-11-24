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
