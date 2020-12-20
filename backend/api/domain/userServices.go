package domain

import (
	"errors"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func GetAllUsers() []models.User {
	return userRepository.GetAllUsers()
}

func UpdateUser(id int, name string) error {
	return userRepository.UpdateUserName(id, name)
}

func UserLogin(userLoginDto dto.UserLoginDto) (error, dto.LoginResponseDto) {
	uerr, user := userRepository.GetUser(userLoginDto.Email, userLoginDto.Password)

	if uerr != nil {
		return uerr, dto.LoginResponseDto{}
	}
	if user.UserID == 0 {
		return errors.New("User not found in database"), dto.LoginResponseDto{}
	}

	var id int
	var canLogin = user.Status
	var err error
	if !canLogin {
		return errors.New("User can't log in because inactive state"), dto.LoginResponseDto{UserID: user.UserID}
	}
	if user.Rol == "employee" {
		err, id = employeeRepository.GetEmployeeId(user.UserID)
		if err != nil {
			return err, dto.LoginResponseDto{}
		}
	}
	if user.Rol == "company" {
		err, id = companyRepository.GetCompanyId(user.UserID)
		if err != nil {
			return err, dto.LoginResponseDto{}
		}
	}
	return err, dto.LoginResponseDto{UserID: user.UserID, SecondaryID: id, Rol: user.Rol}
}

func UpdateUserPassword(userDto dto.UpdateEmployeePasswordDto, userID int) error {
	return userRepository.UpdatePassword(userDto.Password, userID)
}
