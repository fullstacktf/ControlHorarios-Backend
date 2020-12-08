package domain

import (
	"strconv"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(employeeDto dto.CreateEmployeeRequestDto, companyID int) error {
	user := models.User{Username: employeeDto.Username, Password: employeeDto.Password, Email: employeeDto.Email, Rol: employeeDto.Rol}
	id := infrastructure.CreateUser(user)
	employee := models.Employee{UserID: id, CompanyID: companyID, FirstName: employeeDto.FirstName, LastName: employeeDto.LastName}
	return infrastructure.CreateEmployee(employee)
}

func CheckIn(checkInDto dto.CheckInRequestDto, employeeID int) (error, int) {
	record := models.EmployeeRecord{Description: checkInDto.Description, StartTime: time.Now(), EmployeeID: employeeID}
	return infrastructure.CreateRecord(record)
}

func CheckOut(recordID int) error {
	return infrastructure.UpdateRecordTimeOut(recordID, time.Now())
}

func GetSummary(records []models.EmployeeRecord, c *gin.Context) []models.EmployeeRecord {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	infrastructure.DB().
		Select("record_id", "description", "start_time", "end_time", "employee_id").
		Where("employee_id = ?", id).
		Find(&records)
	return records
}

func GetEmployee(id int) models.Employee {
	return infrastructure.GetEmployee(id)
}
