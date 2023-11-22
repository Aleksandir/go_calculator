package main

import "testing"

func TestCalculate(t *testing.T) {
	// Test cases for the calculate function
	tests := []struct {
		num1     int
		num2     int
		operator string
		want     int
	}{
		{1, 2, "+", 3},
		{3, 2, "-", 1},
		{2, 3, "*", 6},
		{6, 3, "/", 2},
		{1, 2, "invalid", 0},
		{1, 0, "/", 0},
		{0, 0, "/", 0},
	}

	for _, tt := range tests {
		got := calculate(tt.num1, tt.num2, tt.operator)
		if got != tt.want {
			t.Errorf("calculate(%v, %v, %q) = %v; want %v", tt.num1, tt.num2, tt.operator, got, tt.want)
		}
	}
}
