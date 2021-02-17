package main

import "fmt"

func main() {
	x := increment()

	x(1)
	x(1)
	x(1)
	fmt.Println(x(1))

}

func increment() func(int) int {
	count := 0
	return func(num int) int {
		count += num
		return count
	}
}
