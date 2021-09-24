package main

import "fmt"

func main() {
	fmt.Println(div(2, 1))
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
