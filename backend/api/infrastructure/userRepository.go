package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func GetUser(email string, password string) models.User {
	var user models.User

	DB().Debug().Where("email = ? AND password = ?", email, password).Find(&user)
	return user
}

func CreateUser(user models.User) int {
	DB().Debug().Create(&user)
	return user.UserID
}
