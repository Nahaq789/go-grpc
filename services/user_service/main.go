package main

import (
	"fmt"
	"log"
	"net"
	"user_service/controller"
	"user_service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := gin.Default()

	// controller := &controller.UserController{}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		listen, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterUserServiceServer(s, &controller.UserController{})

		fmt.Println("Server is running on localhost:50051")
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Println("Server is running on localhost:8080")
	router.Run()
}
