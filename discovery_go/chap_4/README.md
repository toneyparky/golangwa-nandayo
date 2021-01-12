## 1. 이번 장에서 만들어본 계산기 함수는 각 연산자와 숫자 사이에 빈칸이 반드시 하나씩 있어야 한다. 빈칸이 여럿 있더라도 잘 동작하게 개선해보자.

~~~go
package calc

import (
	"strconv"
	"strings"
)

// String set type
type StrSet map[string]struct{}

// NewStrSet returns a new StrSet with the give strs.
func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// BinOp is a binary operator type.
type BinOp func(int, int) int

// PrecMap is a map keyed by operator to set of higher precedence operators.
type PrecMap map[string]StrSet

// NewEvaluator creates a new evaluator with the given binary operations and
// precedence information. The evaluator takes the expression and returns the
// evaluation result.
func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

// Eval returns the evaluation result of the given expr.
// The expr can have operators defined in opMap and
// decimal integers.
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
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
	for _, token := range strings.Split(expr, " ") {
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
	return nums[0]
}
~~~

### 배운점

- strings.Fields()를 이용하면 연속되는 공백을 제거하고 각 요소를 문자열 슬라이스로 반환한다.
