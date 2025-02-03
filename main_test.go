package main

import "testing"

func TestFibonacciUnit(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for _, test := range tests {
		result := fibonacci(test.input)
		if result != test.expected {
			t.Errorf("fibonacci(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
