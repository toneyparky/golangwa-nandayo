package main

import "testing"

func Example_array() {
	Question1([3]string{"사과", "바나나", "파인애플"})
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 파인애플은 맛있다.
}

func TestQuestion1(t *testing.T) {
	Question1([3]string{"사과", "바나나", "파인애플"})
}
