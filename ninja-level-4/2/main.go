package main

import "fmt"

func main() {
	slice := []int{10, 20, 3, 223, 33, 22, 134, 24, 66, 85}

	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", slice)
}
