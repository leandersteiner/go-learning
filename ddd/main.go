package main

import (
	"context"
	"fmt"
	"net"
	"time"

	user2 "ddd/gen/proto/go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	go runServer()

	time.Sleep(time.Second)

	runClient()
}

func runServer() error {
	listenOn := "127.0.0.1:8081"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	user2.RegisterUserServiceServer(server, &userServiceServer{})
	fmt.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type userServiceServer struct {
	user2.UnimplementedUserServiceServer
}

// rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {};

func (u *userServiceServer) CreateUser(ctx context.Context, request *user2.CreateUserRequest) (*user2.CreateUserResponse, error) {
	fmt.Println("created user:", request.User)
	return &user2.CreateUserResponse{}, nil
}

func runClient() error {
	connectTo := "127.0.0.1:8081"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to user service on %s: %w", connectTo, err)
	}

	userService := user2.NewUserServiceClient(conn)
	if _, err := userService.CreateUser(context.Background(), &user2.CreateUserRequest{
		User: &user2.User{
			Id:       1,
			Username: "landi",
			Email:    "leandersteiner@yahoo.de",
		},
	}); err != nil {
		return fmt.Errorf("failed to CreateUser: %w", err)
	}

	fmt.Println("created user")
	return nil
}
