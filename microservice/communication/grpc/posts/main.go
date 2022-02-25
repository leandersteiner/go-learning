package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	comm "learning/microservice/communication/grpc"
	"log"
	"net/http"
	"time"
)

type post struct {
	Id      uint
	Title   string
	Content string
}

func main() {
	mux := chi.NewRouter()
	mux.Post("/posts", func(w http.ResponseWriter, r *http.Request) {

	})

	address := "localhost:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	c := comm.NewCommentsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CommentCreated(ctx, &comm.Comment{
		Id:      1,
		Content: "Test",
		PostId:  1,
	})
	if err != nil {
		log.Fatalln("could not add product:", err)
	}
	log.Println("Product inserted successfully", r)
}
