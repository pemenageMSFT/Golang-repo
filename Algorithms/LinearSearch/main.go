package main

import "fmt"

func LSearch(oss []int, bow int) bool {
	for i := 0; i < len(oss); i++ {
		if oss[i] == bow {
			return true
		}
	}
	return false
}

func main() {
	oss := []int{3, 6, 9, 12, 18}
	// Search for 9 in the int slice
	fmt.Println(LSearch(oss, 9))
}
