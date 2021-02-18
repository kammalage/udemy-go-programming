package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	mUsers, err := json.Marshal(users)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(mUsers)

	var newUsers []user
	err = json.Unmarshal(mUsers, &newUsers)
	fmt.Println(newUsers)

}
