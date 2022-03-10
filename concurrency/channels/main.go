package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go produceInt(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func produceInt(ch chan<- int) {
	defer close(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
}
