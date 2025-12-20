package main

import "fmt"

func main() {
	// 1. 값 타입: 배열 (Array) - 크기가 고정됨 [3]
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 999

	fmt.Println("배열(값) a: ", a) // [1 2 3]

	// 2. 참조 타입: 슬라이스 (Slice) - 크기 미지정 []
	x := []int{1, 2, 3}
	y := x // x가 가리키는 '주소'를 y에게 줌
	y[0] = 999

	fmt.Println("슬라이스(참조) x: ", x) // [999 2 3]
}
