package main

import (
	"fmt"
	"time"
)

func main() {
	var counter Counter
	fmt.Println(counter.String())
	counter.Increment()
	fmt.Println(counter.String())
}

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
}
