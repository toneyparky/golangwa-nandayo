package main

import (
	"fmt"
	"testing"
)

func ExampleQuestion3() {
	answer := Question3([]string{"가", "나", "다", "라"}, "다")
	fmt.Println(answer)
	// Output:
	// 2
}

func TestQuestion3(t *testing.T) {
	Question3([]string{"가", "나", "다", "라"}, "다")
}
