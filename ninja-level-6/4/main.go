package main

import (
	"fmt"
	"strconv"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Shan",
		last:  "Fernando",
		age:   28,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Println(p.first + " " + p.last + "'s" + " age is " + strconv.Itoa(p.age))
	fmt.Println("My age is ", p.age, " just to say it again.")
}
