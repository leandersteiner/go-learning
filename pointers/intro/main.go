package main

import "fmt"

func main() {
	var x int32 = 10
	var y bool = true

	pointerX := &x
	pointerY := &y
	var pointerZ *string

	fmt.Println(pointerX, pointerY, pointerZ)
	fmt.Println(*pointerX, *pointerY)
}
