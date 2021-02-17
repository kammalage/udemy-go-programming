package main

import "fmt"

const (
	_    = iota
	four = iota + 2016
	three
	two
	one
)

func main() {

	fmt.Println(four, three, two, one)
}
