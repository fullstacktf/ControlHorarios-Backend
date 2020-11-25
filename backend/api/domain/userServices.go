package domain

import (
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers() []models.User {
	var users []models.User
	infrastructure.DB.Find(&users)
	return users
}

func CreateUser(user models.UserCompany, c *gin.Context) error {

	result := infrastructure.DB.Debug().
		Select(`User.username,User.email, User.password, User.rol`).Create(&user.User)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})
	return nil
}
