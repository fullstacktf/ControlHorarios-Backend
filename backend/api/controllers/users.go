package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

type UsersRepository struct{}

func GetUsers(c *gin.Context) {
	var users models.Users
	users.Get()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) uint {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return 0
	}

	_, id := user.Nuevo()
	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})
	return id
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	user.Update(id)
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}

func DeleteUser(c *gin.Context) {

	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	user.Delete(id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
