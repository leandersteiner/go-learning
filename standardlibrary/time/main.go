package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	d := 2 * time.Hour + 30 * time.Minute // d is of type time.Duration
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	d2 := time.Now()
	fmt.Println(d2)
	fmt.Println(reflect.TypeOf(d2))

	d3 := d2
	fmt.Println(d2.Equal(d3))

	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))

	after := time.After(2 * time.Second)

	fmt.Println("After:", <-after)

	//tick := time.Tick(2 * time.Second)

	//for {
	//	fmt.Println("Tick:", <-tick)
	//}

	afterFunc := time.AfterFunc(2 * time.Second, func() {
		fmt.Println("I ran")
	})
	defer afterFunc.Stop()

	time.Sleep(4 * time.Second)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		fmt.Println("Ticker:", <-ticker.C)
	}
}
