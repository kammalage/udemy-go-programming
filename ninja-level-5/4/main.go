package main

import "fmt"

func main() {
	shan := struct {
		firstName string
		lastName  string
	}{
		firstName: "Shan",
		lastName:  "Fernando",
	}

	fmt.Println(shan)

	//anonymous function

	func(x int) {
		fmt.Println("We're in an anonymous function", x)
	}(10)
}
