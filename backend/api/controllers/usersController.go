package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/domain"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

type UsersRepository struct{}

func GetUsers(c *gin.Context) {
	users := domain.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
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
	var userLoginDto dto.UserLoginDto
	c.BindJSON(&userLoginDto)
	if userLoginDto.Email == "" || userLoginDto.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	canLogin, userResponseLoginDto := domain.UserLogin(userLoginDto)
	if !canLogin {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User is inactive or Password wrong"})
	} else {
		c.JSON(http.StatusOK, gin.H{"UserID": userResponseLoginDto.UserID, "SecondaryID": userResponseLoginDto.SecondaryID, "Rol": userResponseLoginDto.Rol})
	}
}

func UpdateUserPassword(c *gin.Context) {
	var userDto dto.UpdateEmployeePasswordDto
	err := c.BindJSON(&userDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	DBErr := domain.UpdateUserPassword(userDto, id)

	if DBErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error updating password"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfuly"})
	}
}
