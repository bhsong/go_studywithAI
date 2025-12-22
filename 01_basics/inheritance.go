package main

import "fmt"

type Person struct {
	firstName string
}

func (a *Person) name() string {
	return a.firstName
}

type User struct {
	Person
	lastName string
}

func main() {
	bob := Person{firstName: "Bob"}
	mike := User{Person: Person{firstName: "Mike"}, lastName: "Tyson"}

	fmt.Println(bob.name())
	fmt.Println(mike.firstName + " " + mike.lastName)
}
