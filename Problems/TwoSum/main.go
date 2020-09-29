package main

import "fmt"

// TwoSum solution
func TwoSum(numbers []int, desired int) []int {
	m := make(map[int]int)
	for index, number := range numbers {
		if value, found := m[desired-number]; found {
			return []int{value, index}
		}
		m[number] = index
	}
	return nil
}

func main() {
	numbers := []int{3, 6, 9, 12, 300}
	oss := TwoSum(numbers, 9)
	if oss != nil {
		fmt.Print("The following index add up to 9: ", oss)
	} else {
		fmt.Println("Could not perform two sum for the number 9 :'(")
	}
}
