package main

import "fmt"

func main() {

	test := outerFunc()

	test()

}

func outerFunc() func() {
	fmt.Println("outer func")
	return func() {
		fmt.Println("sum func")
	}
}
