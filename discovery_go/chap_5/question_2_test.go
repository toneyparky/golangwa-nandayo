package main

import (
	"fmt"
	"github.com/jaeyeom/gogo/interface/nocase"
	"sort"
	"testing"
)

// 기존의 test 메서드
func ExampleCaseInsensitive_sort() {
	apple := nocase.CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}

func TestCaseInsensitiveSort(t *testing.T) {
	cases := []struct {
		in   []string
		want []string
	}{
		{[]string{"iPhone", "iPad", "MacBook", "AppStore"}, []string{"AppStore", "iPad", "iPhone", "MacBook"}},
		{[]string{"iPhone", "Apple", "MacBook", "AppStore"}, []string{"Apple", "AppStore", "iPhone", "MacBook"}},
		{[]string{"iPhone", "baby", "BigMac", "AppStore"}, []string{"AppStore", "BigMac", "baby", "iPhone"}},
	}

	for _, c := range cases {
		got := nocase.CaseInsensitive(c.in)
		if compare(got, c.want) {
			t.Errorf("CaseInsensitive(%s) == %s, wand %s", c.in, got, c.want)
		}
	}
}

func compare(one []string, two []string) bool {
	for i, _ := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
