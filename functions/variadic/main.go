package main

import "fmt"

func main() {
	result := addTo(5, 1, 3, 5)
	fmt.Println(result)
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
