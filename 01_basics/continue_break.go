package main

import "fmt"

func main() {
	slice := []string{"Golang", "Ruby", "Javascript", "Python"}
	for _, value := range slice {
		if value == "Golang" {
			fmt.Println("Golang found")
			continue
		}

		if value == "Javascript" {
			fmt.Println("Javascript found")
			break
		}

		fmt.Println(value, "is not Golang")
	}
}
