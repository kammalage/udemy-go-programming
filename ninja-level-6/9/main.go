package main

import "fmt"

func main() {
	total("This is a test string", func() {
		fmt.Println("Sending in this callback fun()")
	})

}

func total(test string, callback func()) {
	fmt.Println(test)
	callback()
}
