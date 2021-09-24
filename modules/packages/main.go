package main

import (
	"fmt"
	"learning/modules/packages/formatter"
	"learning/modules/packages/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
