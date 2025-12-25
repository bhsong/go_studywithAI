package main

import (
	"fmt"
	"sync"
	"time"
)

func boilWater(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("물 끓이기 시작")
	time.Sleep(time.Second * 3)
	fmt.Println("물 끓이기 완료")
}

func chopOnion(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("양파 썰기 시작")
	time.Sleep(time.Second)
	fmt.Println("양파 썰기 완료")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go boilWater(&wg)
	go chopOnion(&wg)

	wg.Wait()

	fmt.Println("라면 완성! 맛있게 드세요!")
}
