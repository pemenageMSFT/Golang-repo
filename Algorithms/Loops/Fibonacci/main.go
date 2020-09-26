package main

import "fmt"

// Fibonacci - each value is the sum of the previous two values

// BowFibLoo
func BowFibLoo(bow int) int {
	first, second := 0, 1
	for ; bow > 1; bow-- {
		first, second = second, first+second
	}
	return second
}

func main() {
	// Looped Fibonacci Style Bows to Sensei
	oss := BowFibLoo(3)
	fmt.Printf("%v Looped Bows to Sensei Fibonacci style!!!", oss)
}
