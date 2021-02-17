package main

import "fmt"

func main() {

	array := [5]int{0, 1, 2, 3, 4}

	for _, v := range array {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", array)
}
