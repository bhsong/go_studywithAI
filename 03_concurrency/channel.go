package main

import "fmt"

func cook(order chan string) {
	order <- "햄버거"
	order <- "피자"
	order <- "파스타"
	close(order)
}

func main() {
	order := make(chan string)

	go cook(order)

	for food := range order {
		fmt.Println(food + "가 서빙되었습니다")
	}
}
