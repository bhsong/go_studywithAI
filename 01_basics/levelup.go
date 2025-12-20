package main

import "fmt"

func levelUp(level *int) {
	*level = *level + 1
}

func main() {
	var level int = 1

	fmt.Println("Current Level: ", level)
	levelUp(&level)
	fmt.Println("Current Level: ", level)
}
