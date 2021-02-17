## 1. 라우터를 사용하여 API 핸들러를 네 종류(Get, Put, Post, Delete)로 나누었을 때, 코드에 반복되는 패턴이 있다. 4장에서 배운 패턴의 추상화를 활용하여 반복되는 패턴을 추상화해보자. 

### 배운점

~~~go

// 메서드 추출
func pringLogWhenError(err error) {
    if err != nil {
        log.Println(err)
        return
    }
}

// 사용 예시
pringLogWhenError(err)
~~~

- 코드를 여러번 보아도 위와 같이 메서드 추출 하나 밖에 못하겠다.
- apiPostHandler과 apiPutHandler 내부의 tasks를 돌면서 인코딩 하는 부분을 추상화하려했지만 맵을 활용하는 메서드가 각각 put과 post라서 타입이 맞지 않기에 추상화를 할 수 없었다. 
- 조금 더 고민해보고 수정하자. 

---

## 2. 이번 장에서 작성한 코드는 개별 할 일 항목만 볼 수 있게 API가 구성되어 있다. 모든 할 일 목록을 볼 수 있는 API를 작성하고 구현해보자. 또한 HTML로도 모든 할 일의 목록들을 볼 수 있게 구현해보자.  

### 배운점

- taskman.go에 라우팅 추가
~~~go
    s.HandleFunc("/", apiMultiGetHandler).Methods("GET")
~~~

- response.go의 dto에 tasks 추가
~~~go
type Response struct {
	ID    task.ID       `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Tasks []task.Task	`json:"tasks"`
	Error ResponseError `json:"error"`
}
~~~

- handler.go에 apiMultiGetHandler 추가 (기존 메서드 apiSingleGetHandler로 변경)
~~~go
func apiMultiGetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := m.GetAll()

	err = json.NewEncoder(w).Encode(Response{
		Tasks:  tasks,
		Error: ResponseError{err},
	})
	pringLogWhenError(err)
}
~~~

- task.html에 복수 조회 api 사용, 응답 값을 템플릿엔진으로 보여주기
~~~go

// api
xhr.open("GET", "/api/v1/task/", false);
xhr.send();
var response = JSON.parse(xhr.responseText);
var tasks = response.tasks;
xhr.send("tasks="+encodeURIComponent(JSON.stringify(tasks)));

// template
{{range $val := .Tasks}}
  <p value="{{$val}}">{{$val.Title}}</p>
{{end}}
~~~

- map의 전체 값을 가져오는 방법에 대한 논의 [링크](https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map)

---

## 3. 맵에는 쓰기 연산을 동시에 할 수 없다. InMemoryAccessor가 문제가 발생하지 않도록 코드를 수정해보자. 동시성 부분의 내용을 읽어야 이 문제를 해결할 수 있을 것이다.  

### 배운점

- 7장 동시성에서 배운 RWMutex를 활용하여 아래와 같이 수정한다. 

~~~go

type ConcurrentInMemoryAccessor struct {
	accessor InMemoryAccessor
	mutex    *sync.RWMutex
}

func (c ConcurrentInMemoryAccessor) Get(key ID) Task {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.accessor.tasks[key]
}

func (c ConcurrentInMemoryAccessor) Set(key ID, value Task) {
	c.mutex.Lock()
	c.accessor.tasks[key] = value
	c.mutex.Unlock()
}
~~~
