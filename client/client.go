package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/XRS0/Forms/services/auth/gen"

	"google.golang.org/grpc"
)

func main() {
	connAuthService, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewAuthServiceClient(connAuthService)
	resp, err := c.Login(context.Background(), &pb.LoginRequest{Username: "John", Password: "жопажопа"})
	if err != nil {
		log.Fatalf("could not login	: %v", err)
	}
	fmt.Println("Message:", resp.Message)
}
