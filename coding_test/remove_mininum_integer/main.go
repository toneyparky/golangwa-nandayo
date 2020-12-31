package main

import "fmt"

func solution(arr []int) []int {
	min := arr[0]
	for _, each := range arr {
		if min > each {
			min = each
		}
	}

	var result []int

	for _, each := range arr {
		if min != each {
			result = append(result, each)
		}
	}

	if len(result) == 0 {
		result = append(result, -1)
	}

	return result
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4}))
}
