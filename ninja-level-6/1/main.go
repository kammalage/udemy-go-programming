package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 100
}

func bar() (num int, s string) {
	num = 12
	s = "test"
	return
}
