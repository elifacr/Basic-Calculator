package main

import (
	"fmt"
	"os"
)

func main() {
	var number1, number2 float64
	var operation rune

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operation (Options: +, -, *, /): ")
	fmt.Fscanf(os.Stdin, "%c", &operation)

	var result float64

	switch operation {
	case '+':
		result = number1 + number2
	case '-':
		result = number1 - number2
	case '*':
		result = number1 * number2
	case '/':
		if number2 == 0 {
			fmt.Println("Division by zero is not allowed.")
			return
		}
		result = number1 / number2
	default:
		fmt.Println("Invalid operation type.")
		return
	}

	fmt.Printf("%g %c %g = %g", number1, operation, number2, result)
}
