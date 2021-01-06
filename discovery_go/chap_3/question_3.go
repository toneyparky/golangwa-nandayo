package main

import (
	"fmt"
	"sort"
)

func Question3(sortedStrings []string, target string) int {
	return sort.SearchStrings(sortedStrings, target)
}

func main() {
	answer := Question3([]string{"가", "나", "다", "라"}, "다")
	fmt.Println(answer)
}
