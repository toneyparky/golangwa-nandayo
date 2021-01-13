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

---

## 2. 이번 장에서 만들어본 계산기 함수는 에러 처리를 하지 않는다. 정수와 에러를 돌려주도록 프로그램을 개선해보자. 에러가 발생하지 않은 경우에는 에러로 nil 값을 돌려주면 되고, 에러가 발생한 경우에는 정수 결과값으로 0과 nil이 아닌 에러값을 돌려주면 된다.

### 배운점

- 디렉토리 하나에는 하나의 패키지명만 존재할 수 있다.
- 패키지명이 같은 파일들끼리는 변수, 함수, 메서드 명을 공유한다. 다만 대문자로 시작하여 외부에 노출되도록 열어준 경우에만 해당한다.
- 에러 처리를 해봤다. 그렇지만 이렇게 하는게 맞는지는 추후에 공부해보며 보완할 예정.

---

## 3. 중괄호로 둘러싸인 수식과 글이 섞여 있는 문자열이 있다. 여기서 수식 부분을 계산하여 답으로 교체하려고 한다. 예를 들어 아래와 같은 글이 있다. 이를 계산하여 결과를 돌려받는 함수를 작성하라.

> 다들 그 동안 고생이 많았다.
> 첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.
> 둘째는 임야 {10 ** 5 mod 7777}평을 가져라.
> 막내는 {10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고 배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.

> 다들 그 동안 고생이 많았다.
> 첫째는 분당에 있는 48평 아파트를 갖거라.
> 둘째는 임야 6676평을 가져라.
> 막내는 3324평 임야를 갖고 배기량 632cc의 경운기를 갖거라.

- 이번 장에서 작성해 본 계산기 함수를 이용하고 `regexp.ReplaceAllStringFunc*`함수를 이용하면 어렵지 않게 작성할 수 있다. 참고로 입력 문자열은 다음과 같은 방법으로 생성하면 편리하다.

~~~go
in := strings.Join([]string{
    "다들 그 동안 고생이 많았다.",
    "첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
    "둘째는 임야 {10 ** 5 mod 7777}평을 가져라.",
    "막내는 {10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고 배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.",
}, "\n")
~~~

- 정규식은 전역 공간에 두면 한 번만 정규식을 컴파일하므로 시스템 자원을 아낄 수 있다. 역따옴표 문자는 따옴표와 마찬가지로 둘러싸서 문자열을 표현하는데 따옴표는 역슬래시를 특수러치하는 데 비하여 역따옴표는
  둘러싸인 문자 그대로를 이용한다. 일반적으로 따옴표를 많이 쓰지만 정규식과 같이 역슬래시가 자주 나올 수 있는 곳에서는 주로 역따옴표를 쓴다. 아래 정규식은 중괄호까지 포함해서 중괄호로 둘러싸인 부분을 찾는다.

~~~go
func EvalReplaceAll(in string) string {
    rx := regexp.MustComplie(`{[^}]+}`)
    out := rx.ReplaceAllStringFunc(in, func(expr string) string {
    	// TODO
    })  
	fmt.Println(out)
}
~~~

- 추가적인 히늩로는 아래와 같이 expr에서 양쪽에 있는 중괄호와 빈칸을 제거할 수 있다. 중괄호 사이에 빈칸이 하나 있음에 유의하자.

~~~go
expr = strings.Trim(expr, "{ }")
~~~

### 배운점

- 왜인지 README.md에서 복사한 코드는 import가 안되는 경우가 있다. 다시 타이핑하자.
- ReplaceAllStringFunc를 사용할 때에 인자로 익명함수를 넘기는게 재미있다.
- Trim 함수의 사용, 추가적으로 여러 라이브러리를 잘 활용하는 것이 중요하다. 
