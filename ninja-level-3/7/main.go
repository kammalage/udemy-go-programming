package main

import "fmt"

func main() {
	a := 100

	if a > 200 {
		fmt.Println("a is larger than 200")
	} else if a > 50 {
		fmt.Println("a is larger than 50")
	} else {
		fmt.Println("a is less than 50")
	}
}
