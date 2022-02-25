package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("Hello"))
	log.Println("Hi")
}

func main() {
	wait := make(chan struct{})
	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      http.TimeoutHandler(handler{}, 2*time.Second, "Timeout"),
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := srv.Serve(l)
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-wait
}
