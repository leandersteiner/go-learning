package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person

	bob := person{}
	bob.name = "Bob"
	fmt.Println(bob.name)

	julia := person{
		"Julia",
		30,
		"cat",
	}

	beth := person{
		age:  40,
		name: "Beth",
	}

	fmt.Println(fred)
	fmt.Println(bob)
	fmt.Println(julia)
	fmt.Println(beth)
}
