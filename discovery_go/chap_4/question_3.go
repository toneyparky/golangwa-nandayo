package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func EvalReplaceAll(in string) string {
	rx := regexp.MustCompile(`{[^}]+}`)

	out := rx.ReplaceAllStringFunc(in, func(expr string) string {
		expr = strings.Trim(expr, "{ }")
		eval := createEval()
		i := eval(expr)
		return strconv.Itoa(i)
	})

	return out
}

func createEval() func(expr string) int {
	return NewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})
}

func main() {
	in := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 {10 ** 5 mod 7777}평을 가져라.",
		"막내는 {10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고 배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.",
	}, "\n")

	all := EvalReplaceAll(in)
	fmt.Println(all)
}
