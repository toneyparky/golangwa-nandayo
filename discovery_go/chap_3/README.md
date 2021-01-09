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

---

## 3. 정렬된 문자열 슬라이스가 있을 때, 특정 문자열이 슬라이스에 있는지를 조사하는 함수를 작성해보자. 이진 검색 알고리즘을 사용하면 좋다.

### 배운점

- sort 패키지의 Search 함수. (Search<T>를 이용하여 타입을 지정 가능, 이미 정렬이 되어있어야 사용 가능)

---

## 4. 슬라이스를 이용하여 큐 자료구조를 구현해보자. 큐에 자료를 넣는 것은 `append` 함수를 이용하고 자료를 꺼낼 때에는 아래의 코드를 이용해보자. 이 방법을 이용하여 자료를 계속해서 넣고 꺼낼 때, 어떤 문제가 발생할 수 있는지 생각해보자. 혹은 큰 문제가 되지 않을 수 있는 이유 역시 생각해보자.

> q = q[1:]

### 배운점

- 함수의 반환시 복수 개의 값을 넘길 수 있다.

---

## 5. 같은 원소가 여러 번 들어갈 수 있는 집합인 `MultiSet`을 기본적으로 제공하는 맵을 이용하여 만들어보자. 아래와 같은 함수들을 작성하면 된다.

~~~go
// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int

// Insert 함수는 집합에 val을 추가한다.
func Insert(m map [string]int, val string)

// Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는 경우에는 아무 일도 일어나지 않는다.
func Erase(m map[string]int, val string)

// Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func Count(m map[string]int, val string) int

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈 칸으로 구분하여 넣은 문자열을 반환한다.
func Strint(m map[string]int) string
~~~

- 아래와 같은 예제 코드가 동작하면 된다.

~~~go
func ExampleMultiSet() {
m := NewMultiSet()
fmt.Println(String(m))
fmt.Println(Count(m, "3"))
Insert(m, "3")
Insert(m, "3")
Insert(m, "3")
Insert(m, "3")
fmt.Println(String(m))
fmt.Println(Count(m, "3"))
Insert(m, "1")
Insert(m, "2")
Insert(m, "5")
Insert(m, "7")
Erase(m, "3")
Erase(m, "5")
fmt.Println(Count(m, "3"))
fmt.Println(Count(m, "1"))
fmt.Println(Count(m, "2"))
fmt.Println(Count(m, "5"))
// Output:
// { }
// 0
// { 3 3 3 3 }
// 4
// 3
// 1
// 1
// 0
}
~~~

### 배운점

- map 자료구조의 활용
- `bytes.Buffer`를 활용한 concat

---

## 6. 입출력 함수들을 잘 살펴보자. `io.Reader`에는 `Read` 메서드가 제공된다. 그러나 파일을 읽을 때에 이 메서드를 이용하기보다는 `io.ReadFull`을 이용해야 편리한 경우가 많다. `io.ReadFull`을 이용하지 않고 `Read` 메서드를 이용하면 어떤 함정에 빠질 수 있는지 생각해보자.

- [Read 메서드](https://golang.org/pkg/io/#Reader)
- [ReadFull 메서드](https://golang.org/pkg/io/#ReadFull)

### 배운점

- Read 메서드의 설명을 참고해보면 Read를 구현할 때에 0과 nil을 반환하는 경우에 대해서 End Of Line인지 여부를 따로 체크해줘야 한다는 이야기가 있다. 하지만 ReadFull은 정확히 주어진
  버퍼의 길이만큼만 읽고 버퍼 길이 이하만큼 읽었을 때에 발생하는 에러를 처리하면 된다.

---

## 7. 입출력 예제에서 `ReadFrom`과 `WriteTo`가 `*os.File` 대신에 `io.Reader` 및 `io.Writer`를 넘겨받도록 작성될 때 어떤 장점이 생기는지 생각해보자.

- [ReadFrom 메서드](https://golang.org/pkg/bytes/#Buffer.ReadFrom)
- [WriteTo 메서드](https://golang.org/pkg/bytes/#Buffer.WriteTo)

### 배운점

- 추상화와 관련되어 있는 장점이라고 생각한다. [io.Reader](https://golang.org/pkg/io/#Reader)
  , [io.Writer](https://golang.org/pkg/io/#Writer) 는 인터페이스다. 직접 File을 사용하지 않고 필요한 행위만 인터페이스를 통해 사용하면 캡슐화 등등의 추상화로 인한
  긍정적인 장점을 얻을 수 있다. 
