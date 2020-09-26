package main

import "fmt"

// Factorial - the product (multiplication) of an integer and the integers below it

// BowLoop factorial func
func BowLoop(bow int) int {
	var BowFactor = 1
	for ; bow > 1; bow-- {
		BowFactor *= bow
	}
	return BowFactor
}

func main() {
	oss := BowLoop(6)
	fmt.Printf("%v Loop Factorial Bows to Sensei!!!", oss)
}
