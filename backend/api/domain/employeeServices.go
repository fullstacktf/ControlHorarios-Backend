package domain

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(employee models.UserEmployee, c *gin.Context, id int) error {

	employee.Employee.UserID = id
	employee.Employee.CompanyID, _ = strconv.Atoi(c.Params.ByName("idCompany"))
	result := infrastructure.DB().Debug().
		Select(`Employee.first_name, Employee.last_name`).Create(&employee.Employee)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New employee created successfully"})
	return nil
}

func CreateCheckIn(employeeRecord models.EmployeeRecord, c *gin.Context) error {

	employeeRecord.EmployeeID, _ = strconv.Atoi(c.Params.ByName("idCompany"))

	result := infrastructure.DB.Debug().Create(&employeeRecord)

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
	result := infrastructure.DB.Debug().
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

	id, _ := strconv.Atoi(c.Params.ByName("idCompany"))

	result := infrastructure.DB.Debug().
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
