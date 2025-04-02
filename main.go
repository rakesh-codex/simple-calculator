package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number1> <operation> <number2>")
		fmt.Println("Operations: add, sub, mul, div")
		fmt.Println("Example: go run main.go 5 add 3")
		os.Exit(1)
	}

	// Extract arguments
	num1Str := os.Args[1]
	operation := os.Args[2]
	num2Str := os.Args[3]

	// Convert string arguments to float64 for flexibility
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	if err1 != nil {
		fmt.Println("Error: First argument must be a number")
		os.Exit(1)
	}

	num2, err2 := strconv.ParseFloat(num2Str, 64)
	if err2 != nil {
		fmt.Println("Error: Second argument must be a number")
		os.Exit(1)
	}

	// Perform the operation
	var result float64
	switch operation {
	case "add":
		result = num1 + num2
	case "sub":
		result = num1 - num2
	case "mul":
		result = num1 * num2
	case "div":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operation. Use add, sub, mul, or div")
		os.Exit(1)
	}

	// Display the result
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operation, num2, result)
}
