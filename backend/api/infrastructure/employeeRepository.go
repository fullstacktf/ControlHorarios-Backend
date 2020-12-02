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
