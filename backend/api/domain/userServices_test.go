package domain

import (
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/stretchr/testify/assert"
)

type UserRepositoryFake struct{}

func (u UserRepositoryFake) GetUser(email, password string) (error, models.User) {
	return nil, models.User{UserID: 1, Rol: "employee", Status: true}
}

func (u UserRepositoryFake) CreateUser(user models.User) (error, int) {
	return nil, user.UserID
}

func (u UserRepositoryFake) UpdatePassword(password string, id int) error {
	return nil
}

func (u UserRepositoryFake) UpdateUserName(id int, name string) error {
	return nil
}

func (u UserRepositoryFake) GetAllUsers() []models.User {
	var users []models.User
	return users
}

type CompanyRepositoryFake struct{}

func (c CompanyRepositoryFake) GetCompanyId(id int) (error, int) {
	return nil, 0
}

func (c CompanyRepositoryFake) CreateCompany(company models.Company) (error, int) {
	return nil, company.CompanyID
}

func (c CompanyRepositoryFake) GetCompany(id int) models.Company {
	var company models.Company
	return company
}

type EmployeeRepositoryFake struct{}

func (e EmployeeRepositoryFake) GetEmployeeId(id int) (error, int) {
	employee := models.Employee{EmployeeID: 1}
	return nil, employee.EmployeeID
}

func (e EmployeeRepositoryFake) GetEmployee(id int) models.Employee {
	var employee models.Employee
	return employee
}

func (e EmployeeRepositoryFake) CreateEmployee(employee models.Employee) error {
	return nil
}

func (e EmployeeRepositoryFake) GetEmployeesByCompanyID(id int) []models.Employee {
	var employees []models.Employee
	return employees
}

func (e EmployeeRepositoryFake) UpdateEmployee(id int, status bool) error {
	return nil
}

func TestUserLoginShouldReturnUserIDs(t *testing.T) {
	SetRepositories(UserRepositoryFake{}, EmployeeRepositoryFake{}, CompanyRepositoryFake{})
	_, responseDto := UserLogin(dto.UserLoginDto{})
	assert.Equal(t, dto.LoginResponseDto{UserID: 1, SecondaryID: 1, Rol: "employee"}, responseDto)
}
