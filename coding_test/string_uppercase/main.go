package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func solution(s string) string {
	splited := strings.Split(s, " ")

	var answer []string

	for _, word := range splited {
		answer = append(answer, changeCase(word))
	}

	return strings.Join(answer, " ")
}

func changeCase(word string) string {
	var b bytes.Buffer

	for idx, char := range word {
		if idx%2 == 1 {
			b.WriteRune(unicode.ToLower(char))
		} else {
			b.WriteRune(unicode.ToUpper(char))
		}
	}
	return b.String()
}

func main() {
	fmt.Println(solution("asdfasdfasdf asdf d asdf fdsa"))
}
