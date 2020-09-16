package main

import "fmt"

type Sensei struct {
	Firstname string
	Lastname  string
	Age       int
	Proverbs  []string
}

type HigherBeing interface {
	bow()
}

func (s *Sensei) bow() {
	fmt.Print("BOW TO YOUR SENSEI:", s.Firstname, " ", s.Lastname, "!!!")
}

func bowToYourSensei(s *Sensei) {
	s.bow()
}

func main() {
	s1 := Sensei{
		Firstname: "Professor",
		Lastname:  "Sensei",
		Age:       30,
		Proverbs:  []string{"I am all knowing!", "I hope someone laughs at this", "BOW TO YOUR SENSEI!"},
	}
	// one way to bow
	bowToYourSensei(&s1)
	fmt.Println("\n")
	// another way to pay homage to sensei all mighty
	s1.bow()

}
