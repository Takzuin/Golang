package main

import (
	"fmt"
)

func main() {
	// Declare variables for the two numbers and the operator
	var num1, num2 float64
	var operator string

	// Read the first number from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	// Read the second number from the user
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Read the operator from the user
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Perform the operation based on the entered operator
	switch operator {
	case "+":
		// Add the two numbers and print the result
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		// Subtract the second number from the first and print the result
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		// Multiply the two numbers and print the result
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		// Divide the first number by the second and print the result if the second number is not zero
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			// Print an error message if the second number is zero
			fmt.Println("Error: Division by zero")
		}
	default:
		// Print an error message if the entered operator is not valid
		fmt.Println("Invalid operator")
	}
}
