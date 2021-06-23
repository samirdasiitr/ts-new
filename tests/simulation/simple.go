package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ts-new/model"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	req := &pb.NewUserRequest{
		Name:         "testuser",
		LastName:     "testlastname",
		MobileNumber: "9812345678",
		Email:        "test@test.com",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, req)
	if err != nil {
		log.Fatalf("could not register new: %v", err)
	}

	log.Printf("User: %s", r)
}
