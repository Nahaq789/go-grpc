package main

import (
	"fmt"
	"log"
	"organization_service/controller"
	"organization_service/model"
	"organization_service/proto"
	"organization_service/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initClient() proto.UserServiceClient {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	client := proto.NewUserServiceClient(conn)
	log.Println("Connected to user service")
	return client
}

func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userServiceClient := initClient()

	c := controller.NewOrganizationController(
		repository.NewOrganizationRepository(initDB(), userServiceClient),
	)
	router.POST("/organization", c.CreateOrganization)
	fmt.Println("Organization service is running...")

	router.Run(":9090")
}

func initDB() *gorm.DB {
	dsn := "docker:docker@tcp(localhost:3307)/test_database_organization?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	log.Println("Database connection established")

	err = db.AutoMigrate(&model.Organization{})
	if err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}
	return db
}
