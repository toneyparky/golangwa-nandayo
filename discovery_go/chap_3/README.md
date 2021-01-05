## 1. 배열 예제에서 사용된 다음 코드는 공교롭게도 모두 마지막 글자가 받침이 없는 과일들뿐이었습니다. 마지막 글자에 받침이 있는 경우에도 어색하지 않게 조사가 붙어서 출력되도록 다음 코드를 수정하시오.

~~~go
func Example_array() {
fruits := [3]string{"사과", "바나나", "토마토"}
for _, fruit := range fruits {
fmt.Printf("%s는 맛있다.\n", fruit)
}
// Output:
// 사과는 맛있다
// 바나나는 맛있다
// 토마토는 맛있다.
}
~~~

### 배운점

- _test를 붙여서 테스트 파일 만들기
- Example을 이용한 테스트 (따로 실행해야하는데 보완이 필요)
- rune (go에서는 한글을 저장할 때에 3byte로 저장한다. 그리고 그 형식은 utf-8를 따른다. 이와 같은 문자를 저장할 때에 rune을 사용하고 이 rune은 int32의 별칭이다.)
- unicode를 사용할 경우 한글의 받침 유무를 알기 위해서는 28로 나눈 나머지를 확인하면 된다.

---

## 2. []int 슬라이스를 넘겨받아서 오름차순으로 정렬하는 함수를 작성해보자. 슬라이스 a의 i 번째 값과 j 번째 값을 서로 맞바꿀 때는 아래의 코드를 참고하자.

> a[i], a[j] = a[j], a[i]

### 배운점

- 두개의 변수 할당 ([여러 변수](https://stackoverflow.com/questions/21071507/can-you-declare-multiple-variables-at-once-in-go)
  , [타입이 다른 변수들](https://stackoverflow.com/questions/45086082/multiple-variables-of-different-types-in-one-line-in-go-without-short-variable))
- 버블 정렬
