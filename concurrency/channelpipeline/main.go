package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		time.Sleep(100 * time.Millisecond)
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for {
		x, ok := <-in
		if !ok {
			break
		}
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for {
		num, ok := <-in
		if !ok {
			fmt.Println("Work done")
			break
		}
		fmt.Println(num)
	}
}
