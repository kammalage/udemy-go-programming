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

	jaeger := person{
		firstName:    "Eren",
		lastName:     "Jaeger",
		faveIcecream: []string{"Vanilla"},
	}

	//fmt.Println(shan, jaeger)

	m := map[string]person{
		"Fernando": shan,
		"Jaeger":   jaeger,
	}

	m2 := map[string]person{
		shan.lastName: shan,
	}

	for i, v := range m["Fernando"].faveIcecream {
		fmt.Println(i, v)
	}

	for i, v := range m["Jaeger"].faveIcecream {
		fmt.Println(i, v)
	}

	fmt.Println(m2)
}
