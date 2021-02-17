package main

import "fmt"

func main() {
	const (
		a = 10
		b = 10
		c = 20
		d = 30
	)

	if a == b {
		fmt.Println("numbers equal each other")
	} else {
		fmt.Println("numbers do not equal each other")
	}

	if a <= c {
		fmt.Println("first number is less than second number")
	} else {
		fmt.Println("second number is greater than or equal to first number")
	}

	if a >= b {
		fmt.Println("first number is greater than or equal to second number")
	} else {
		fmt.Println("second number is greater than or equal to first number")
	}

	if c > b {
		fmt.Println("first number is greater than second number")
	} else {
		fmt.Println("second number is greater than first number")
	}

	if b < c {
		fmt.Print("first number is less than second number")
	} else {
		fmt.Print("first number is greater than second number")
	}

}
