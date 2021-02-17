package main

import "fmt"

func main() {
	xs := []string{"James", "Bond", "shaken not stirred"}
	ys := []string{"Money", "Penny", "helloooo james bond"}

	agents := [][]string{xs, ys}

	fmt.Println(agents)
}
