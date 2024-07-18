package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	// Test for a normal case
	result, err := Factorial(5)
	expected := 120
	if err != nil || result != expected {
		t.Errorf("Expected %v, got %v (error: %v)", expected, result, err)
	}

	// Test for an edge case
	result, err = Factorial(0)
	expected = 1
	if err != nil || result != expected {
		t.Errorf("Expected %v, got %v (error: %v)", expected, result, err)
	}

	// Test for a negative input
	result, err = Factorial(-1)
	if err == nil {
		t.Errorf("Expected an error for negative input, got %v", result)
	}

	// Test for another normal case
	result, err = Factorial(7)
	expected = 5040
	if err != nil || result != expected {
		t.Errorf("Expected %v, got %v (error: %v)", expected, result, err)
	}
}
