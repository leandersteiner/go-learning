package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	numerator := 20
	denominator := 3

	remainder, err := calcRemainder(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder)
}

func calcRemainder(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return numerator / denominator, nil
}
