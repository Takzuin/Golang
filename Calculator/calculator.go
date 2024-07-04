package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Leer números del usuario
	fmt.Print("Ingresa el primer número: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scanln(&num2)

	// Leer el operador del usuario
	fmt.Print("Ingresa un operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Realizar la operación
	switch operator {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: División por cero")
		}
	default:
		fmt.Println("Operador no válido")
	}
}
