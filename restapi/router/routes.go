package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pbJWT "github.com/XRS0/Forms/services/JWT/gen"
	pbAuth "github.com/XRS0/Forms/services/auth/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// AuthCredentials структура для хранения данных аутентификации
type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool `json:"success"`
}

func RouteAuth(r *gin.Engine) {

	authGroup := r.Group("/auth")

	authGroup.POST("/jwt", func(c *gin.Context) {
		var creds AuthCredentials
		if err := c.BindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tokenString, err := CreateToken(creds.Username, ConnectToJWTService())
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"jwt": tokenString, // Здесь должна быть генерация и отправка настоящего JWT
		})
	})

	authGroup.POST("/credCheck", func(c *gin.Context) {
		var creds AuthCredentials
		if err := c.BindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		success, err := AuthUser(ConnectToAuthService(), creds.Username, creds.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// success := creds.Username == "sum" && creds.Password == "jwt"

		c.JSON(http.StatusOK, gin.H{"success": success})
	})
}

func ConnectToAuthService() pbAuth.AuthServiceClient {
	connAuth, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil
	}
	cAuth := pbAuth.NewAuthServiceClient(connAuth)
	return cAuth
}

func ConnectToJWTService() pbJWT.JWTServiceClient {
	connJWT, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	cJWT := pbJWT.NewJWTServiceClient(connJWT)
	return cJWT
}

func AuthUser(cAuth pbAuth.AuthServiceClient, username, password string) (*pbAuth.LoginResponse, error) {
	respAuth, errAuth := cAuth.Login(context.Background(), &pbAuth.LoginRequest{Username: username, Password: password})
	if errAuth != nil {
		fmt.Printf("could not login (authModule): %v\n", errAuth)
		return nil, errAuth
	}
	return respAuth, nil
}

func CreateToken(username string, cJWT pbJWT.JWTServiceClient) (*pbJWT.CreateTokenResponse, error) {
	respJWT, errJWT := cJWT.CreateJWTToken(context.Background(), &pbJWT.CreateTokenRequest{Username: "John"})
	if errJWT != nil {
		return nil, grpc.Errorf(7, "че то пошло не по плану")
	}
	tokenString := respJWT.Token
	fmt.Println("Token:", respJWT.Token)
	return &pbJWT.CreateTokenResponse{Token: tokenString}, nil
}
