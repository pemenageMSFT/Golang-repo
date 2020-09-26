package main

import "fmt"

// Fibonacci - each value is the sum of the previous two values

// BowFib
func BowFib(bow int) int {
	if bow < 2 {
		return bow
	}
	return BowFib(bow-1) + BowFib(bow-2)
}

func main() {
	// Bow to Sensei Fibonacci style
	oss := BowFib(9)
	fmt.Printf("%v Bows Recursive Fibonacci style!!!", oss)
}
