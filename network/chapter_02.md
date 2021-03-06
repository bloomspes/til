# 무엇을 공부 했는가?

네트워크의 `1계층`, `2계층`의 동작 원리에 대해서 학습한다.

- 네트워크 회선의 구조로 부터 일어날 수 있는 전기 신호의 원리를 통해 1계층에서 신호를 데이터로 교환하는 `원리`에 대해 알 수 있었다.

- 신호를 데이터로 교환해서 각각의 수신처로 전달하는 경우에 생길 수 있는 충돌과 간섭을 막기 위한 방법들을 찾게 되고 `규칙`을 만들게 된다. 2계층의 핵심은 데이터를 받기 위한 규칙이라고 요약할 수 있겠다.

- 마지막으로, `스위치`를 언급하면서 네트워크 물리 3계층에 대해서도 살짝 개념 확장이 들어간다. 

- 특정 포트의 패킷을 수신하기 위한 스위치는 어디서 사용할까? 클라우드 컴퓨팅 환경을 사용하면서 L4, L7 스위치에 대해서 많이 들어본 경험이 있을 수 있겠다.

- 서버의 로드밸런싱에 들어가는 이 스위치는 특정 수신처의 정보를 입력받게 되어 서버의 부하를 분산시키는 역할을 한다.

## 궁금한 점

1. 신호의 충돌은 허브에서 일어나는데, 허브는 각각의 프레임을 구분할 수 있는 테이블이 존재하지 않는 것일까? 

- 학습을 할 수 있으면 방향을 정해서 목적지 까지 신호를 흘려보낸다.

- 허브의 동작 원리는 학습할 수 있는 테이블이 존재하지 않으므로 전기 신호를 플러딩(흘려보내는 것)과 같다.

2. 스위치는 어드레스 테이블을 `학습` 하여 필터링을 한다고 한다. 만일 학습된 테이블을 초기화 한다면 어떤 주기로 진행이 되는 것일까? 프리엠블과 관련이 있을까?

## 피드백

- 기존의 문제점을 언급하며 그의 보완재로 나온 것임을 인과 관계를 밝혀서 설명하자.