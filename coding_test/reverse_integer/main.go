package main

import "fmt"

func solution(n int64) []int {
	var ints []int

	for n != 0 {
		ints = append(ints, int(n%10))
		n = n / 10
	}

	return ints
}

func main() {
	fmt.Println(solution(12345))
}
