package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("작업 중... ", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	// 함수 앞에 go 키워드를 붙이
	go sayHello(&wg)
	wg.Wait()

	fmt.Println("메인 함수 종료")
}
