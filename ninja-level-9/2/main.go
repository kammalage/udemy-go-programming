package main

import "fmt"

type human interface {
	speak()
}

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

	saySomething(&p)

}

func (p *person) speak() {
	fmt.Println("My name is", (*p).first, (*p).last, "and I am ", (*p).age, " years old.")
}

func saySomething(h human) {
	h.speak()
}
