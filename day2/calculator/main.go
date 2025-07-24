package main

import "fmt"

func main() {
	var operator string
	fmt.Println("We can do add, multilpy, subtract, divide.")
	fmt.Println("Enter the operator: +, -, *, / ")
	fmt.Scanln(&operator)

	var num1, num2 float64
	fmt.Println("Enter First number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second number: ")
	fmt.Scanln(&num2)

	// if operator == "+" {
	// 	fmt.Println("The addition of two numbers is :", num1+num2)
	// } else if operator == "-" {
	// 	fmt.Println("The Substraction of two numbers is :", num1-num2)
	// } else if operator == "*" {
	// 	fmt.Println("The Multiplication of two numbers is :", num1*num2)
	// } else if operator == "/" {
	// 	fmt.Println("The Division of two numbers is :", num1/num2)
	// } else {
	// 	fmt.Println("invalid operator")
	// }

	switch operator {
	case "+":
		fmt.Println("The addition of two numbers is :", num1+num2)
	case "-":
		fmt.Println("The Substraction of two numbers is :", num1-num2)
	case "*":
		fmt.Println("The Multiplication of two numbers is :", num1*num2)
	case "/":
		fmt.Println("The Division of two numbers is :", num1/num2)
	default:
		fmt.Println("invalid operator")
	}
}
