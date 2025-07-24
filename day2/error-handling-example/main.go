package main

import (
	"errors"
	"fmt"
)

// return new error if n is -ve
// errors.New is the function from errors package that creates the new error value with given message
func checkNumber(n int) error {
	if n < 0 {
		return errors.New("negative numbers are not allowed")
	}
	return nil
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	//call the checknumber func and store in err
	err := checkNumber(n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Valid number entered: ", n)
}
