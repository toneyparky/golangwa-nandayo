package main

import "fmt"

func main() {
	var count int
	fmt.Println("몇 회 반복할까요?")
	_, err := fmt.Scanln(&count)

	if err != nil {
		panic(err)
	}

	for i := 1; i <= count; i++ {
		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", i*10, (i+1)*10)
	}
}
