package main

import "fmt"

var result int

func multiply(a int, b int) int {
	return a * b
}

func main() {
	result = multiply(5, 6)
	fmt.Println("Multiplication of 5 and 6 is", result)
}
