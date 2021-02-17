package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{10, 20, 30}

	x = append(x, y...)

	fmt.Println(x)
}
