package main

import "fmt"

func main() {
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)

	const size = 10
	var arr [size]int

	var y = [...]int{1, 2, 3}
	var z = [3]int{1, 2, 3}
	fmt.Println(y == z) // prints true
	fmt.Println(arr)
}
