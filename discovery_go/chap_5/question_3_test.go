package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExampleID_UnmarshalJSON() {
	myStruct := MyStruct{
		"title",
		"hello world",
		22,
		NewID("9223372036854775807"),
	}
	value, err := json.Marshal(myStruct)

	if err != nil {
		log.Println(err)
		return
	}

	ms := MyStruct{}
	err = json.Unmarshal(value, &ms)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(ms.Title)
	fmt.Println(ms.Internal)
	fmt.Println(ms.ID.value)
	// Output:
	// title
	//
	// 9223372036854775807
}
