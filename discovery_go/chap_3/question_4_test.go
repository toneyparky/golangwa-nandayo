package main

import "testing"

func ExampleQuestion4() {
	Question4()
	// Output:
	// Enqueued:  [1]
	// Enqueued:  [1 2]
	// Enqueued:  [1 2 3]
	// 1
	// 2
	// 3
}

func TestQuestion4(t *testing.T) {
	Question4()
}
