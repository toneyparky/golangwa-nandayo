package main

import (
	"bytes"
)

type MultiSet map[string]int

func NewMultiSet() MultiSet {
	return MultiSet{}
}

func Insert(m MultiSet, input string) {
	m[input]++
}

func Erase(m MultiSet, input string) {
	m[input]--
}

func Count(m MultiSet, input string) int {
	return m[input]
}

func String(m MultiSet) string {
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
