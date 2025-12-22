package main

import "fmt"

func main() {
	var arr [3]int = [3]int{10, 20, 30}

	fmt.Println("배열:", arr)

	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5)

	fmt.Println("슬라이스:", slice)
	fmt.Println("슬라이스 길이:", len(slice))

}
