package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIcecream []string
}

func main() {
	shan := person{
		firstName:        "Shan",
		lastName:         "Fernando",
		favoriteIcecream: []string{"Vanilla", "Cookies and Cream"},
	}

	stranger := person{
		firstName:        "Stranger",
		lastName:         "Danger",
		favoriteIcecream: []string{"???", "???"},
	}

	fmt.Println(shan.firstName)
	fmt.Println(shan.lastName)
	for _, v := range shan.favoriteIcecream {
		fmt.Println("\t", v)
	}

	fmt.Println(stranger.firstName)
	fmt.Println(stranger.lastName)
	for _, v := range stranger.favoriteIcecream {
		fmt.Println("\t", v)
	}
}
