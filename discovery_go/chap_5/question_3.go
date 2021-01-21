package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type MyStruct struct {
	Title    string `json:"title"`
	Internal string `json:"-"'`
	Value    int64  `json:",omitempty"`
	ID       *ID
}

type ID struct {
	value string
}

func NewID(value string) *ID {
	return &ID{value}
}

func (id ID) String() string {
	return id.value
}

func (id ID) MarshalJSON() ([]byte, error) {
	parseInt, err := strconv.ParseInt(id.value, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return strconv.AppendInt(nil, parseInt, 10), nil
}

func (id *ID) UnmarshalJSON(data []byte) (err error) {
	s2 := string(data)
	id.value = s2
	return nil
}

func main() {
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
}
