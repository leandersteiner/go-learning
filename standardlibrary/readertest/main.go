package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "Mein Name ist Leander Steiner"
	sr := strings.NewReader(s)

	count, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func countLetters(r io.Reader) (int, error) {
	buf := make([]byte, 16)
	result := 0

	for {
		n, err := r.Read(buf)
		for i, b := range buf[:n] {
			fmt.Println("Index", i, ":", string(b))
			result++
		}
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return 0, err
		}
		fmt.Println("Durchgang fertig. Count:", result)
	}
}
