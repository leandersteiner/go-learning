package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	name, err := getName(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your name is " + name)
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press enter when done.\n"
	_, _ = fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}
