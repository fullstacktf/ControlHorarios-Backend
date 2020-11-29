package domain

import (
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers() []models.User {
	var users []models.User
	infrastructure.DB().Find(&users)
	return users
}

func CreateUser(user models.UserCompany, c *gin.Context) (error, int) {

	result := infrastructure.DB().Debug().
		Select(`User.username,User.email, User.password, User.rol`).Create(&user.User)
	if result.Error != nil {
		return result.Error, 0
	}

	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})
	return nil, user.User.UserID
}

func CreateUserEmployee(user models.UserEmployee, c *gin.Context) (error, int) {

	result := infrastructure.DB().Debug().
		Select(`User.username,User.email, User.password, User.rol`).Create(&user.User)
	if result.Error != nil {
		return result.Error, 0
	}

	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})
	return nil, user.User.UserID
}

func UpdateUser(id int, name string) error {
	result := infrastructure.DB().Debug().
		Model(&models.User{}).
		Where("users.user_id = ?", id).
		Updates(models.User{
			Username: name,
		})
	return result.Error
}

func DeleteUser(id int) error {
	result := infrastructure.DB().Debug().Delete(&models.User{}, id)
	return result.Error
}
