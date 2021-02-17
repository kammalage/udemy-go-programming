package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	sum := func(s []int) int {
		var total int
		for _, value := range s {
			total += value
		}
		return total
	}(s)

	fmt.Println(sum)
}
