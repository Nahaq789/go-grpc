package main

import (
	"user_service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller := &controller.UserController{}
	router.GET("/ping", controller.CreateUser)

	router.Run()
}
