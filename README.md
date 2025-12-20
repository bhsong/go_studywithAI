# 🐹 Go Study with AI: Golang Study Repository

**AI(Artificial Intelligence)와 함께 Go 언어를 학습하고 예제 코드를 기록하는 저장소입니다.**

이곳에는 기초 문법부터 알고리즘, 미니 프로젝트까지 AI의 도움을 받아 작성하고 분석한 코드들이 저장됩니다.

## 🎯 학습 목표 (Goals)
- Go 언어의 핵심 문법과 철학(Go way) 이해하기
- 동시성 프로그래밍(Goroutines, Channels) 마스터하기
- AI를 활용하여 효율적으로 코드 작성 및 리팩토링 연습하기
- 매일/매주 꾸준히 1커밋 실천하기

## 📂 폴더 구조 (Directory Structure)
```bash
.
├── 01_basics          # 변수, 제어문, 함수 등 기초 문법
├── 02_data_structures # 배열, 슬라이스, 맵, 구조체
├── 03_concurrency     # 고루틴, 채널 등 동시성 예제
├── 04_algorithms      # 알고리즘 문제 풀이
└── playground         # 자유롭게 테스트한 코드

## 💡 적용 방법
[2025-12-20] 
- Go 포인터는 연산 불가 (p++, p + 1 같은 거 안됨)
- 지역 변수의 주소 반환 가능 (알아서 스택이 아닌 힙에 저장)
- free() 안해도 가비지컬렉터가 알아서 메모리 해제
- 초기값 nil
- nil은 의미상으로는 Null과 같지만 '0'은 아니며, '0'과 비교 불가능
- nil은 Zero Value 개념으로 참조 타입 변수(포인터, 슬라이스 ,맵, 인터페이스)의 기본값은 nil 
- 포인터 쓸 때 nil 인지 확인 필요
- Array(배열) = 값 타입(Value Type), Slice = 참조 타입(Reference Type)
