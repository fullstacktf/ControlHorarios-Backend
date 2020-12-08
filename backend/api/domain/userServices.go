package domain

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func GetAllUsers() []models.User {
	return infrastructure.GetAllUsers()
}

func UpdateUser(id int, name string) error {
	return infrastructure.UpdateUserName(id, name)
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

func UpdateUserPassword(userDto dto.UpdateEmployeePasswordDto, userID int) error {
	return infrastructure.UpdatePassword(userDto.Password, userID)
}
