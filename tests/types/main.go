package main

import "fmt"

type Event string

const (
	Created Event = "created"
	Deleted Event = "deleted"
)

func main() {
	str := asString()
	evt := asEvent()

	fmt.Println(Event(str) == evt, str == string(evt))

	switch evt {
	case Created:
		fmt.Println(evt)
	case Deleted:
		fmt.Println(evt)
	}

	switch Event(str) {
	case Created:
		fmt.Println(str)
	case Deleted:
		fmt.Println(str)
	}
}

func asString() string {
	return "deleted"
}

func asEvent() Event {
	return Created
}
