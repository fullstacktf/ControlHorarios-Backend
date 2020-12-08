package domain

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func GetAllUsers() []models.User {
	var users []models.User
	infrastructure.DB().Find(&users)
	return users
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

func UserLogin(userLoginDto dto.UserLoginDto) dto.LoginResponseDto {
	user := infrastructure.GetUser(userLoginDto.Email, userLoginDto.Password)
	var id int
	if user.Rol == "employee" {
		id = infrastructure.GetEmployeeId(user.UserID)
	}
	if user.Rol == "company" {
		id = infrastructure.GetCompanyId(user.UserID)
	}
	return dto.LoginResponseDto{UserID: user.UserID, SecondaryID: id, Rol: user.Rol}
}
