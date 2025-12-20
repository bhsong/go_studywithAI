package main

import "fmt"

func calc(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	val1, val2 := calc(10, 5)
	fmt.Println("합:", val1)
	fmt.Println("차:", val2)
}
