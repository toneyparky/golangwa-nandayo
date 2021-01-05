package main

import "fmt"

func Question2(input []int) []int {
	length := len(input)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func main() {
	answer := Question2([]int{4, 2, 1, 6, 3, 5})

	for _, number := range answer {
		fmt.Println(number)
	}
}
