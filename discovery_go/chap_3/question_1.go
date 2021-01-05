package main

import "fmt"

func Question1(words [3]string) {
	for _, word := range words {
		lastRune := getLastRune(word)
		if hasSyllableFinalSound(lastRune) {
			fmt.Printf("%s은 맛있다.\n", word)
			continue
		}
		fmt.Printf("%s는 맛있다.\n", word)
	}
}

func hasSyllableFinalSound(character rune) bool {
	return character%28 != 16
}

func getLastRune(word string) rune {
	return []rune(word)[len(word)/3-1]
}

func main() {
	Question1([3]string{"사과", "바나나", "파인애플"})
}
