package main

import "fmt"

func main() {
	leander := Person{
		FirstName: "Leander",
		LastName:  "Steiner",
		Age:       24,
	}

	fmt.Println(leander.String())
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
