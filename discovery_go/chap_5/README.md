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
