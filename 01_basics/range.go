package main

import "fmt"

var slice = []string{"Golang", "Ruby", "Javascript", "Python"}

var mapping = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
}

func main() {
	for index, value := range slice {
		fmt.Println(index, value)
	}

	for key, value := range mapping {
		fmt.Println(key, value)
	}
}
