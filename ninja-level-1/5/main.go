package main

import "fmt"

type exampleint int //int is underlying type of exampleint

var y int

func main() {
	var x exampleint
	fmt.Printf("%v\t%T\t\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v\t%T\t", y, y)
}
