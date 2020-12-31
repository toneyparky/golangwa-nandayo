package main

import "fmt"

func main() {
	Hanoi(3)
}

func Hanoi(count int) {
	fmt.Printf("Number of disks: %d\n", count)
	playHanoi(count, 1, 3, 2)
}

func playHanoi(count int, one int, two int, three int) {
	if count == 1 {
		moveHanoi(one, three)
	} else {
		playHanoi(count-1, one, three, two)
		moveHanoi(one, three)
		playHanoi(count-1, two, one, three)
	}
}

func moveHanoi(from int, to int) {
	fmt.Printf("%d -> %d\n", from, to)
}
