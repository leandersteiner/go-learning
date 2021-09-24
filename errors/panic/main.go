package main

import "os"

func main() {
	doPanic(os.Args[0])
}

func doPanic(msg string) {
	panic(msg)
}
