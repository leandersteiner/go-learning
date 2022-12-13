package main

import (
	"fmt"
)

func main() {
	svc := &Service{host: "localhost"}
	err := svc.Start()
	if err != nil {
		panic(err)
	}

	svc2 := &Service{}
	svc2.host = "google.com"
	svc2.Start()

	var n number = 5
	println(n.square())
	println(n.cube())
}

type Service struct {
	host string
}

func (s *Service) Start() error {
	fmt.Println("connecting to", s.host)
	return nil
}

type number int

func (n number) square() number {
	return n * n
}

func (n number) cube() number {
	return n.square() * n
}

type writer interface {
	Write(p []byte) (n int, err error)
}

func write(w writer, c string) {}

type resource struct {
  path string
  data []byte
}

func (r *resource) deny() {
  fmt.Println("access denied")
}

type image struct {
  r resource
}
