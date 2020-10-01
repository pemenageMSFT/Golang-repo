package main

import "fmt"

// Reverse a string
func LoopReverseMeSensei(str string) {
	if len(str) == 1 {
		fmt.Print(str)
		return
	}
	LoopReverseMeSensei(str[1:])
	fmt.Print(str[0:1])
}

func main() {
	LoopReverseMeSensei("!!!iesneS ruoy ot woB")
}
