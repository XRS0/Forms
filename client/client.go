package main

import (
	"context"
	"fmt"
	"log"

	pbJWT "github.com/XRS0/Forms/services/JWT/gen"
	pbAuth "github.com/XRS0/Forms/services/auth/gen"
	"google.golang.org/grpc"
)

func main() {
	connAuth, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	connJWT, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cAuth := pbAuth.NewAuthServiceClient(connAuth)
	cJWT := pbJWT.NewJWTServiceClient(connJWT)
	respAuth, errAuth := cAuth.Login(context.Background(), &pbAuth.LoginRequest{Username: "John", Password: "жопажопа"})
	if errAuth != nil {
		log.Fatalf("could not login	: %v", err)
	}
	respJWT, errJWT := cJWT.CreateJWTToken(context.Background(), &pbJWT.CreateTokenRequest{Username: "John", Password: "жопажопа"})
	if errJWT != nil {
		log.Fatalf("could not login	: %v", err)
	}
	fmt.Println("Message:", respAuth.Message)
	fmt.Println("Token:", respJWT.Token)
}
