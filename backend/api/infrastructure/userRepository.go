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

func UpdatePassword(password string, id int) error {
	result := DB().Debug().
		Model(&models.User{}).
		Where("user_id = ?", id).
		Update("password", password)

	return result.Error
}

func UpdateUserName(id int, name string) error {
	result := DB().Debug().
		Model(&models.User{}).
		Where("users.user_id = ?", id).
		Updates(models.User{
			Username: name,
		})
	return result.Error
}

func GetAllUsers() []models.User {
	var users []models.User
	DB().Find(&users)
	return users
}
