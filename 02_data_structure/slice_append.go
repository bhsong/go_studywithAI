package main

import "fmt"

func main() {
	slice := []string{"Golang", "Java"}
	newSlice := append(slice, "Ruby")

	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	fmt.Println(slice, len(slice), cap(slice))
}
