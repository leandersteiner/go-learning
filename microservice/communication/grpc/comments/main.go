package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	comm "learning/microservice/communication/grpc"
	"log"
	"net"
	"net/http"
)

type comments = []comment

var commentStore = make(comments, 0)

type CommentServer struct {
	comm.UnimplementedCommentsServer
}

func (CommentServer) CommentCreated(ctx context.Context, c *comm.Comment) (*comm.Comment, error) {
	response := &comm.Comment{
		Id:      1,
		Content: "Test",
		PostId:  1,
	}
	commentStore = append(commentStore, comment{
		Id:      uint(response.Id),
		Content: response.Content,
		PostId:  uint(response.PostId),
	})
	fmt.Println("comment added", response, commentStore[0])
	return response, nil
}

func (CommentServer) CommentDeleted(ctx context.Context, c *comm.Comment) (*comm.Comment, error) {
	response := &comm.Comment{
		Id:      1,
		Content: "Test",
		PostId:  1,
	}
	return response, nil
}

type comment struct {
	Id      uint
	Content string
	PostId  uint
}

func main() {
	mux := chi.NewRouter()
	mux.Post("/comments", func(w http.ResponseWriter, r *http.Request) {

	})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	s := grpc.NewServer()
	comm.RegisterCommentsServer(s, &CommentServer{})

	log.Println("Starting gRPC listener on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
