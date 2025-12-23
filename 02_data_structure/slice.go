package main

import "fmt"

func main() {
	arr := [...]string{"Golang", "Java"}
	slice := arr[0:2]

	slice[0] = "Ruby"
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

}
