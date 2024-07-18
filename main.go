package main

import "fmt"

func Factorial(n int) (int, error) {
	if n < 0 {
		return -1, fmt.Errorf("factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func main() {
	// Example usage
	n := 5
	result, err := Factorial(n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Factorial of %d is %d\n", n, result)
	}
}
