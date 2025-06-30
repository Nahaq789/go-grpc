package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Logic to create a user
	c.JSON(201, gin.H{"message": "User created"})
}
