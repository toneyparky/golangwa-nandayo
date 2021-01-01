package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	p, q := 0, 1
	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, strconv.Itoa(p))
		p, q = q, p+q
	}
	fmt.Printf(strings.Join(result, ", "))
}
