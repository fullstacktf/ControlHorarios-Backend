package controllers

import (
	"log"
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
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

func EmployeeLogin(c *gin.Context) {
	var employeeLoginDto dto.EmployeeLoginDto
	c.BindJSON(&employeeLoginDto)
	if employeeLoginDto.Email == "" || employeeLoginDto.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	user := domain.EmployeeLogin(employeeLoginDto)
	if user.UserID != 0 {
		c.JSON(http.StatusOK, gin.H{"UserID": user.UserID, "Rol": user.Rol})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email or Password wrong"})
	}
}
