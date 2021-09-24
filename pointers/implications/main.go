package main

import "fmt"

func main() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}
