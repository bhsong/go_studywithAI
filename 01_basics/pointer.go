package main

import "fmt"

// 1. 값에 의한 전달 (복사본))
func changeValue(val int) {
	val = 100 // 복사본을 100으로 바꿈
}

// 2. 포인터에 의한 전달 (주소) -> 매개변수 타입 앞에 * 붙임
func changePointer(ptr *int) {
	*ptr = 200 // 주소(ptr)를 찾아가서 그 안의 값(*ptr)을 100으로 바꿈
}

func main() {
	a := 10

	fmt.Println("변경 전 a: ", a)

	changeValue(a)
	fmt.Println("changeValue 호출 후 a: ", a) // 값에 의한 전달이므로 a는 변경되지 않음

	changePointer(&a)
	fmt.Println("changePointer 호출 후 a: ", a) // 포인터에 의한 전달이므로 a가 변경됨
}
