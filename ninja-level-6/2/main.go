package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5}

	secondNum := []int{2, 4, 3, 5, 4, 3, 4}

	fmt.Println(foo(num...))

	fmt.Println(bar(secondNum))
}

func foo(x ...int) (total int) {
	for _, num := range x {
		total += num
	}
	return
}

func bar(num []int) (total int) {
	for i := 0; i < len(num); i++ {
		total += i
	}
	return
}
