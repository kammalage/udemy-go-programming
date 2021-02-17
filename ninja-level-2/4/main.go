package main

import "fmt"

func main() {
	var ten int = 10

	fmt.Printf("%d\t%b\t%X\n", ten, ten, ten)

	shiftedTen := ten << 1

	fmt.Printf("%d\t%b\t%X", shiftedTen, shiftedTen, shiftedTen)
}
