package main

import "fmt"

func main() {
	mapEx := make(map[string]int)
	mapEx["ele1"] = 1
	fmt.Println(mapEx)

	mapEx["newEle"] = 2
	fmt.Println(mapEx)
	fmt.Println(mapEx["ele1"])
	fmt.Println(mapEx["newEle"])

	delete(mapEx, "ele1")
	fmt.Println(mapEx)
}
