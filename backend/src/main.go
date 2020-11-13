package main

import (
	"../api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var userServices models.UsersServices
	r.GET("/users", userServices.GetUsers)

	r.Run(":8989")
}
