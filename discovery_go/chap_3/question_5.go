package main

import (
	"bytes"
)

func NewMultiSet() map[string]int {
	return map[string]int{}
}

func Insert(m map[string]int, input string) {
	m[input]++
}

func Erase(m map[string]int, input string) {
	m[input]--
}

func Count(m map[string]int, input string) int {
	return m[input]
}

func String(m map[string]int) string {
	var buffer bytes.Buffer
	buffer.WriteString("{ ")

	for key, _ := range m {
		for i := 0; i < m[key]; i++ {
			buffer.WriteString(key)
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("}")

	return buffer.String()
}
