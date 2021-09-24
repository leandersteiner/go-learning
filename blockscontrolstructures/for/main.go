package main

import "fmt"

func main() {

	// Complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Condition-only for statement
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// Infinite for statement
	for {
		fmt.Println("Hello")
	}
}
