package main

import (
	"fmt"
	"log"
	"net"
	"user_service/controller"
	"user_service/model"
	"user_service/proto"
	"user_service/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	// controller := &controller.UserController{}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db := initDB()
	userRepo := repository.NewUserRepository(db)

	userController := controller.NewUserController(userRepo)

	go func() {
		listen, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterUserServiceServer(s, userController)

		fmt.Println("Server is running on localhost:50051")
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Println("Server is running on localhost:8080")
	router.Run()
}

func initDB() *gorm.DB {
	dsn := "docker:docker@tcp(localhost:3308)/test_database_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	log.Println("Database connection established")

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}
	return db
}
