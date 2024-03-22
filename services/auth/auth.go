package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/XRS0/Forms/services/auth/gen"
	"google.golang.org/grpc"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &authServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *authServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("Received: чел с ником", in.Username, "и с паролем", in.Password)
	if in.Username == "sum" && in.Password == "jwt" {
		return &pb.LoginResponse{Message: "туалет открыт"}, nil
	}
	return nil, grpc.Errorf(7, "вы не обладаете полномочиями какать здесь")
}
