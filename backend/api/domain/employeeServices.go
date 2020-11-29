package domain

import (
	"net/http"
	"strconv"

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