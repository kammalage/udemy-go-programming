package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(func(s []int) int {
		var total int
		for _, value := range s {
			total += value
		}
		return total
	}(s))

}
