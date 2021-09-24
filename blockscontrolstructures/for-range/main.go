package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// Ignore(hide) index
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// Ignore value
	for i := range evenVals {
		fmt.Println(i)
	}

	// Iterating over maps
	testMap := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for k, v := range testMap {
		fmt.Println(k, v)
	}

	// Iteration order varies
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	fmt.Println()

	// Iterating over strings
	samples := []string{"hello", "apple_π!"}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}
