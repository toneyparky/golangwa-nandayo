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
