package main

import "fmt"

func main() {
	slice := []string{"Golang", "Java"}
	var slice2 []string

	slice2 = slice
	slice[0] = "Ruby"

	if slice == nil {
		fmt.Println("Slice is nil")
	}

	fmt.Println(slice2)
	fmt.Println(slice)
}
