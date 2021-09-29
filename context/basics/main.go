package main

import (
	"context"
	"fmt"
	"net/http"
)

func useContext(ctx context.Context, data string) (string, error) {
	return "", nil
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(rw.Write([]byte("Hello!")))
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// wrap the context with stuff
		req = req.WithContext(ctx)
		fmt.Println(ctx)
		handler.ServeHTTP(rw, req)
	})
}

func main() {
	handler := http.HandlerFunc(testHandler)
	s := http.Server{
		Addr:    ":8080",
		Handler: Middleware(handler),
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
