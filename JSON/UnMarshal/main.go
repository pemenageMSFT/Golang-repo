package main

import (
	"encoding/json"
	"fmt"
)

type Sensei struct {
	Firstname string
	Lastname  string
	Age       int
	Sayings   []string
}

func main() {
	sensei1 := Sensei{
		Firstname: "Professor",
		Lastname:  "Sensei",
		Age:       30,
		Sayings:   []string{"Hello, I am sensei", "A choke for a choke and the whole world is asleep", "I like turtles"},
	}

	s1, err := json.Marshal(sensei1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON VERSION", string(s1))

	s2 := Sensei{}

	err = json.Unmarshal(s1, &s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Back to normal\n", s2)
}
