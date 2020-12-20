package infrastructure

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

type EmployeeRepositoryInterface interface {
	GetEmployeeId(id int) (error, int)
	GetEmployee(id int) models.Employee
	CreateEmployee(employee models.Employee) error
	GetEmployeesByCompanyID(id int) []models.Employee
	UpdateEmployee(id int, status bool) error
}

type EmployeeRepository struct{}

func (e EmployeeRepository) GetEmployeeId(id int) (error, int) {
	var employee models.Employee
	result := DB().Debug().Where("user_id = ?", id).First(&employee)
	return result.Error, employee.EmployeeID
}

func (e EmployeeRepository) GetEmployee(id int) models.Employee {
	var employee models.Employee
	DB().Debug().
		Joins("User").
		Joins("Company").
		Where("employee_id = ?", id).First(&employee)
	return employee
}

func (e EmployeeRepository) CreateEmployee(employee models.Employee) error {
	result := DB().Debug().Create(&employee)
	return result.Error
}

func (e EmployeeRepository) GetEmployeesByCompanyID(id int) []models.Employee {
	var employees []models.Employee
	DB().Debug().
		Joins("User").
		Joins("Company").
		Where("employee.company_id = ?", id).
		Find(&employees)
	return employees
}

func (e EmployeeRepository) UpdateEmployee(id int, status bool) error {
	result := DB().Debug().Model(&models.User{}).Where("user_id = ?", id).Update("status", status)
	return result.Error
}
