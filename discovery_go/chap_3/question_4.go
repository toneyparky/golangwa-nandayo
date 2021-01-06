package main

import "fmt"

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	fmt.Println("Enqueued: ", queue)
	return queue
}

func dequeue(queue []int) ([]int, int) {
	element := queue[0]
	queue = queue[1:]
	return queue, element
}

func Question4() {
	var queue []int
	var element int
	queue = enqueue(queue, 1)
	queue = enqueue(queue, 2)
	queue = enqueue(queue, 3)

	queue, element = dequeue(queue)
	fmt.Println(element)
	queue, element = dequeue(queue)
	fmt.Println(element)
	queue, element = dequeue(queue)
	fmt.Println(element)
}

func main() {
	Question4()
}
