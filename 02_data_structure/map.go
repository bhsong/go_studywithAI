package main

import "fmt"

func main() {
	//mapEx := make(map[string]string, 2)
	//fmt.Println(mapEx)

	//mapEx["firstName"] = "Mike"
	//mapEx["lastName"] = "Smith"

	var mapEx = map[string]string{
		"firstName": "Mike",
		"lastName":  "Smith",
	}

	fmt.Println(mapEx)
}
