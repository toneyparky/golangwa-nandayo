package main

import "fmt"

// 4장의 생성기 패턴과 5장의 인터페이스를 이용하여 각각 반복자를 구현해보자.

func NewIteratingGenerator(elements []int) func() int {
	current := 0
	length := len(elements)
	return func() int {
		if current >= length {
			panic("Iterated over element length")
		}
		current++
		return elements[current-1]
	}
}

type IteratingInterface interface {
	HasNext() bool
	Next() int
}

type Iterator []int

func (i Iterator) HasNext() bool {
	return len(i) != 0
}

func (i *Iterator) Next() int {
	iterator := *i
	element := iterator[0]
	*i = iterator[1:]
	return element
}

func main() {
	generator := NewIteratingGenerator([]int{1, 2, 3, 4, 5})
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
	//fmt.Println(generator())

	iterator := Iterator([]int{1, 2, 3, 4, 5})

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
