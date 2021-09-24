package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (p Person) processName() {
	p.FirstName = "Erik"
	p.LastName = "Müller"
}

func (p *Person) processNamePtr() {
	p.FirstName = "Erik"
	p.LastName = "Müller"
}

func main() {
	name := "Leander"
	fmt.Println(name)
	fmt.Println(&name)
	processName(&name)

	leander := &Person{
		FirstName: "Leander",
		LastName:  "Steiner",
	}

	fmt.Println(leander)
	leander.processName()
	fmt.Println(leander)
	leander.processNamePtr()
	fmt.Println(leander)

}

func processName(name *string) {
	fmt.Println(name)
	fmt.Println(*name)
}
