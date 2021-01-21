## 1. 이번 장에서 만든 Task 구조체에 MarkDone 메서드를 구현해보자. 이 메서드가 호출되면 해당 작업 및 SubTask들 모두 상태가 DONE으로 바뀐다.

### 배운점

- 배열 내부의 값을 바꿀
  때에는 [포인터를 사용하지 말고 인덱스를 사용하자.](https://stackoverflow.com/questions/20185511/range-references-instead-values) (시간을 엄청
  허비했다.)

~~~go
// task 패키지에 추가한 코드.  
func (t *Task) MarkDone() {
t.Status = DONE

if t.SubTasks != nil {
for idx, task := range t.SubTasks {
t.SubTasks[idx].Status = DONE
task.MarkDone()
}
}
}
~~~

- 외부 패키지의 타입 대해 메서드를 추가하고 싶을
  때에는 [그 타입을 내장하는 다른 타입을 만드는 방법도](https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go)
  있다. (처음에 적용했다가 지금의 방법으로 복귀)

---

## 2. 정렬 인터페이스에 예제로 나와 있는 ExampleCaseInsensitiveSort는 한 가지 경우에만 테스트한다. 테이블 기반 테스트를 이용한 TestCaseInsensitiveSort 함수를 구현하여 여러가지 경우의 수에 대하여 정렬이 제대로 동작하는지 확인해보자.

### 배운점

- case는 예약어이다.
- golang에서는 []string 두 개에 대한 동등성 비교 조차도 직접 만들어야 한다.
- 테이블 기반 테스트의 활용은 어렵지 않다.

---

## 3. JSON으로 구조체를 직렬화하여 보관하고 이것을 역직렬화하여 이용하는 코드가 있다. 원래 필드에 int64형에 이름을 붙인 ID라는 자료형을 이용하였는데 자바스크립트에서 53비트 이상의 정수를 제대로 읽을 수 없는 문제 때문에 JSON 출력에서는 string 형으로 변경하려고 한다. JSON 태그 \`json:",string"\`을 이용하면 되는 문제지만 기존에 저장된 자료를 읽을 때 발생하는 문제가 있어서, 예전 형식인 정수로 되어 있는 경우와 문자열로 되어 있는 경우 모두 읽을 수 있는 역직렬화 코드를 작성하고 싶다. UnmarshalJSON과 형 단언을 이용하여 이것을 가능하게 해보자.

### 배운점

- 문제의 정답을 내지는 못했다고 생각한다. 어떻게 코드를 짜야할지 아직도 감이 잘 안온다. 추후에 다시 보자.
- String(), UnmarshalJSON(), MarshalJSON()을 구현했다. struct를 field로 가질 경우 이 메서드들을 작성해줘야 직렬화, 역직렬화가 원활하게 진행된다.
