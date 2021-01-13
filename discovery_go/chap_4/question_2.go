package main

import (
	"strconv"
	"strings"
)

// String set type
type strSet map[string]struct{}

// newStrSet returns a new strSet with the give strs.
func newStrSet(strs ...string) strSet {
	m := strSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// binOp is a binary operator type.
type binOp func(int, int) int

// precMap is a map keyed by operator to set of higher precedence operators.
type precMap map[string]strSet

// NewEvaluator creates a new evaluator with the given binary operations and
// precedence information. The evaluator takes the expression and returns the
// evaluation result.
func NewEvaluator2(opMap map[string]binOp, prec precMap) func(expr string) (int, error) {
	return func(expr string) (int, error) {
		return Eval2(opMap, prec, expr)
	}
}

// Eval2 returns the evaluation result of the given expr.
// The expr can have operators defined in opMap and
// decimal integers.
func Eval2(opMap map[string]binOp, prec precMap, expr string) (result int, err error) {
	defer func() {
		if err != nil {
			result = 0
		}
	}()

	ops := []string{"("} // 초기 여는 괄호
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Fields(expr) {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산을 하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	result = nums[0]
	return result, err
}
