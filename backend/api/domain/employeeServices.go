package domain

import (
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func CreateEmployee(employeeDto dto.CreateEmployeeRequestDto, companyID int) error {
	user := models.User{Username: employeeDto.Username, Password: employeeDto.Password, Email: employeeDto.Email, Rol: employeeDto.Rol, JoinedDate: time.Now()}
	employee := models.Employee{User: user, CompanyID: companyID, FirstName: employeeDto.FirstName, LastName: employeeDto.LastName}
	return infrastructure.CreateEmployee(employee)
}

func CheckIn(checkInDto dto.CheckInRequestDto, employeeID int) (error, int) {
	record := models.EmployeeRecord{Description: checkInDto.Description, StartTime: time.Now(), EmployeeID: employeeID}
	return infrastructure.CreateRecord(record)
}

func CheckOut(recordID int) error {
	return infrastructure.UpdateRecordTimeOut(recordID, time.Now())
}

func GetSummary(employeeID int) []models.EmployeeRecord {
	return infrastructure.GetRecordsByID(employeeID)
}

func GetEmployee(id int) models.Employee {
	return infrastructure.GetEmployee(id)
}
