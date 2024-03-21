package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/XRS0/Forms/services/JWT/gen"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

type JWTserver struct {
	pb.UnimplementedJWTServiceServer
}

var secretKey = []byte("shopapopa")

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJWTServiceServer(s, &JWTserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *JWTserver) Login(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем набор утверждений (claims)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = in.Username
	claims["password"] = in.Password // Обычно пароль не включают в токен
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Подписываем токен нашим секретным ключом
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenResponse{Token: tokenString}, nil
}

func (s *JWTserver) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи, который мы ожидаем, соответствует алгоритму в заголовке токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["password"])
		return &pb.ValidateTokenResponse{Access: true}, nil
	} else {
		return nil, err
	}
}
