package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	faveIcecream []string
}

func main() {

	shan := person{
		firstName:    "Shan",
		lastName:     "Fernando",
		faveIcecream: []string{"Oreo", "Vanilla"},
	}

	michelle := person{
		firstName:    "Michelle",
		lastName:     "Fernando",
		faveIcecream: []string{"Vanilla"},
	}

	fmt.Println(shan, michelle)

	for i, v := range shan.faveIcecream {
		fmt.Println(i, v)
	}
	for i, v := range michelle.faveIcecream {
		fmt.Println(i, v)
	}
}
