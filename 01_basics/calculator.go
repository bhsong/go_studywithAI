package main

import "fmt"

func calculator(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			fmt.Println("0으로 나눌 수 없습니다")
			return 0
		}
	default:
		fmt.Println("유효하지 않은 연산자입니다")
		return 0
	}
}

func main() {
	var a, b, result int
	var op string

	fmt.Print("첫번째 인수를 입력하세요: ")
	fmt.Scan(&a)
	fmt.Print("두번째 인수를 입력하세요: ")
	fmt.Scan(&b)
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scan(&op)

	result = calculator(a, b, op)
	fmt.Println("결과: ", result)
}
