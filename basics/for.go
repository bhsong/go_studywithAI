package main

import "fmt"

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("성인입니다.")
	} else {
		fmt.Println("미성년자입니다.")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("반복 횟수:", i)
	}

	n := 0
	for n < 3 {
		fmt.Println("while처럼 동작:", n)
		n++
	}
}
