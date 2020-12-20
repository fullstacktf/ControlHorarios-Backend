package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

type UserRepositoryInterface interface {
	GetUser(email, password string) (error, models.User)
	CreateUser(user models.User) (error, int)
	UpdatePassword(password string, id int) error
	UpdateUserName(id int, name string) error
	GetAllUsers() []models.User
}

type UserRepository struct{}

func (u UserRepository) GetUser(email, password string) (error, models.User) {
	var user models.User
	result := DB().Debug().Where("email = ? AND password = ?", email, password).Find(&user)
	return result.Error, user
}

func (u UserRepository) CreateUser(user models.User) (error, int) {
	result := DB().Debug().Create(&user)
	return result.Error, user.UserID
}

func (u UserRepository) UpdatePassword(password string, id int) error {
	result := DB().Debug().
		Model(&models.User{}).
		Where("user_id = ?", id).
		Update("password", password)
	return result.Error
}

func (u UserRepository) UpdateUserName(id int, name string) error {
	result := DB().Debug().
		Model(&models.User{}).
		Where("users.user_id = ?", id).
		Updates(models.User{
			Username: name,
		})

	return result.Error
}

func (u UserRepository) GetAllUsers() []models.User {
	var users []models.User
	DB().Find(&users)
	return users
}
