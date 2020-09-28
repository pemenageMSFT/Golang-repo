package main

import "fmt"

// Grab the last value of a slice [len(<slice_name>)-1]
// Reverse a string
func RecursiveReverseMeSensei(bow string) string {
	if len(bow) == 1 {
		return bow
	}
	return RecursiveReverseMeSensei(bow[1:]) + bow[0:1]
}

func main() {
	// Bow to your Sensei!!!
	oss := RecursiveReverseMeSensei("!!!iesneS ruoy ot woB")
	fmt.Print(oss)
}
