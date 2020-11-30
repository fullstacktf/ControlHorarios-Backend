package controllers

import (
	"log"
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/domain"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EmployeeRepository struct {
}

func CreateEmployee(c *gin.Context) {
	var userEmployee models.UserEmployee

	err := c.ShouldBindWith(&userEmployee, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	_, id := domain.CreateUserEmployee(userEmployee, c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	domain.CreateEmployee(userEmployee, c, id)
}

func CreateCheckIn(c *gin.Context) {
	var employeeRecord models.EmployeeRecord

	err := c.ShouldBindWith(&employeeRecord, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	domain.CreateCheckIn(employeeRecord, c)
}

func DoCheckOut(c *gin.Context) {
	domain.DoCheckOut(c)
}

func UpdateEmployeePassword(c *gin.Context) {
	var user models.User

	err := c.ShouldBindWith(&user, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	domain.UpdateEmployeePassword(user, c)
}
