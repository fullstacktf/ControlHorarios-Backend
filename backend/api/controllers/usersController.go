package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/domain"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

type UsersRepository struct{}

func GetUsers(c *gin.Context) {
	users := domain.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var user models.UserCompany
	err := c.ShouldBindWith(&user, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)

	}
	domain.CreateUser(user, c)
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

	dbErr := domain.UpdateUser(id, user.Username)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	dbErr := domain.DeleteUser(id)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

func UserLogin(c *gin.Context) {
	var employeeLoginDto dto.UserLoginDto
	c.BindJSON(&employeeLoginDto)
	if employeeLoginDto.Email == "" || employeeLoginDto.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	userLoginDto := domain.UserLogin(employeeLoginDto)
	if userLoginDto.UserID != 0 {
		c.JSON(http.StatusOK, gin.H{"UserID": userLoginDto.UserID, "SecondaryID": userLoginDto.SecondaryID, "Rol": userLoginDto.Rol})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email or Password wrong"})
	}
}
