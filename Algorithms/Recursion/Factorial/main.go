package main

import "fmt"

// Factorial- a product (multiplication) of an integer and all the integers below it

// Bow to sensei Function
func RecursiveBow(bow int) int {
	if bow < 2 {
		return bow
	}
	return bow * RecursiveBow(bow-1)
}

func main() {
	a := RecursiveBow(6)
	fmt.Printf("%v Bows to Sensei", a)
}
