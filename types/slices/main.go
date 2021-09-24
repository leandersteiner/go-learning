package main

import "fmt"

func main() {
	var x []int
	x = append(x, 10)
	x = append(x, 4, 5, 6)

	fmt.Println(x)

	y := []int{20, 30, 40}
	x = append(x, y...)

	fmt.Println(x)
	fmt.Println(cap(x))
	x = append(x, 1, 2)
	fmt.Println(cap(x))

	a := []int{1, 2, 3, 4}
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	e := a[:]
	f := a[0:3]

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
}
