package main

import "fmt"

func ExampleIterator() {
	generator := NewIteratingGenerator([]int{1, 2, 3, 4, 5})
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())

	iterator := Iterator([]int{1, 2, 3, 4, 5})

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 1
	// 2
	// 3
	// 4
	// 5
}
