package main

import (
	"fmt"
	"testing"
)

func Example_sort() {
	answer := Question2([]int{4, 2, 1, 6, 3, 5})

	for _, number := range answer {
		fmt.Println(number)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func TestQuestion2(t *testing.T) {
	Question2([]int{4, 2, 1, 6, 3, 5})
}
