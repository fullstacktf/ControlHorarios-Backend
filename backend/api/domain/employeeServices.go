package domain

import (
	"net/http"
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

func CreateCheckIn(employeeRecord models.EmployeeRecord, c *gin.Context) error {

	employeeRecord.EmployeeID, _ = strconv.Atoi(c.Params.ByName("id"))
	employeeRecord.StartTime = time.Now()
	result := infrastructure.DB().Debug().Create(&employeeRecord)

	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "New Record created successfully",
		"recordID": employeeRecord.RecordID})
	return nil

}

func DoCheckOut(c *gin.Context) error {

	timeOut := time.Now()
	recordId, _ := strconv.Atoi(c.Params.ByName("idRecord"))
	result := infrastructure.DB().Debug().
		Model(models.EmployeeRecord{}).
		Where("employee_records.record_id = ?", recordId).
		Updates(models.EmployeeRecord{
			EndTime: timeOut,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateEmployeePassword(user models.User, c *gin.Context) error {

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	result := infrastructure.DB().Debug().
		Model(&models.User{}).
		Where("users.user_id = ?", id).
		Updates(models.User{
			Password: user.Password,
		})

	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated password"})
	return nil
}

func GetSummary(records []models.EmployeeRecord, c *gin.Context) []models.EmployeeRecord {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	println(id)
	infrastructure.DB().
		Select("record_id", "description", "start_time", "end_time", "employee_id").
		Where("employee_id = ?", id).
		Find(&records)
	return records
}

func GetEmployee(id int) models.Employee {
	return infrastructure.GetEmployee(id)
}
