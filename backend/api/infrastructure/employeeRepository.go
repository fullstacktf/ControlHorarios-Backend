package infrastructure

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func GetEmployeeId(id int) int {
	var employee models.Employee
	DB().Debug().Where("user_id = ?", id).First(&employee)
	return employee.EmployeeID
}

func GetEmployee(id int) models.Employee {
	var employee models.Employee
	DB().Debug().Where("employee_id = ?", id).First(&employee)
	return employee
}

func CreateEmployee(employee models.Employee) error {
	result := DB().Debug().Create(&employee)
	return result.Error
}

func GetEmployeesByCompanyID(id int) []models.Employee {
	var employees []models.Employee
	DB().Debug().
		Joins("User").
		Joins("Company").
		Where("employee.company_id = ?", id).
		Find(&employees)
	return employees
}


func UpdateEmployee(UserID int ) error {
    result := DB().Debug().Model(&models.User{}).Where("user_id = ?", UserID).Update("password", "Inactive")
    return result.Error
}