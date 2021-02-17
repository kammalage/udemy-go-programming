package main

import "fmt"

type exampleint int

func main() {
	var x exampleint
	fmt.Printf("%v\t%T\t\n", x, x)
	x = 42
	fmt.Println(x)
}
