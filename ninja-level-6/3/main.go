package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("FUNCTION FOO()")
}

func bar() {
	fmt.Println("FUNCTION BAR()")
}
