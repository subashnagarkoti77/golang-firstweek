package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println("You entered zero.")
	} else if n%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

}
