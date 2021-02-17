package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		{
			fmt.Println("I like to skii!")
		}
	case "swimming":
		{
			fmt.Println("I like swimming")
		}
	case "surfing":
		{
			fmt.Println("I like surfing")
		}
	}
}
