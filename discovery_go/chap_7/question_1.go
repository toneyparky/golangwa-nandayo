package main

import "fmt"

func Quicksort(s []int, low, high int) {
	if low >= high {
		return
	}
	if len(s) <= 1 {
		return
	}
	pivotIdx := partition(s, low, high)
	Quicksort(s, low, pivotIdx-1)
	Quicksort(s, pivotIdx+1, high)
	return
}

func partition(s []int, low, high int) int {
	i := low - 1
	for j := low; j < high; j++ {
		if s[j] < s[high] {
			i++
			s[j], s[i] = s[i], s[j]
		}
	}
	s[i+1], s[high] = s[high], s[i+1]
	return i + 1
}

func main() {
	ints := []int{1, 2, 45, 61, 7, 3, 5, 6, 8}
	Quicksort(ints, 0, len(ints)-1)

	for _, each := range ints {
		fmt.Println(each)
	}
}
