package main

import "fmt"

func main() {
	a := 10
	switch {
	case a > 9:
		{
			fmt.Println("a is greater than 9")
		}
	case a > 3:
		{
			fmt.Println("a is greater than 3")
		}
	case a > 1:
		{
			fmt.Println("a is greater than 1")
		}
	}
}
