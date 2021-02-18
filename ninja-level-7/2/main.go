package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p := &person{
		first: "Shan",
		last:  "Fernando",
		age:   28,
	}

	//fmt.Println(*p)
	p.changeMe()

}

func (p *person) changeMe() {
	fmt.Println(p)
	fmt.Println(*p)
	(*p).first = "Imposter"
	p.last = "Fake Last Name"
	(*p).age = 100
	fmt.Println(p)
	fmt.Println(*p)
}
