package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)

	go createInt(data)

	result := <- data
	fmt.Println(result)
}

func createInt(intCh chan int) {
	time.Sleep(5 * time.Second)
	intCh <- 5
}
