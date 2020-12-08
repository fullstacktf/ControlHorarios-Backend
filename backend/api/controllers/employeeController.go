package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/domain"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EmployeeRepository struct {
}

func CreateEmployee(c *gin.Context) {
	var employeeDto dto.CreateEmployeeRequestDto
	err := c.BindJSON(&employeeDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}

	companyID, _ := strconv.Atoi(c.Params.ByName("id"))
	DBError := domain.CreateEmployee(employeeDto, companyID)

	if DBError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating employee"})
	} else {
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
	}
}

func CreateCheckIn(c *gin.Context) {
	var checkInDto dto.CheckInRequestDto
	err := c.BindJSON(&checkInDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}

	employeeID, _ := strconv.Atoi(c.Params.ByName("id"))
	DBError, checkInID := domain.CreateCheckIn(checkInDto, employeeID)
	if DBError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error in check in"})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message":  "New Record created successfully",
			"recordID": checkInID})
	}
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

func GetSummary(c *gin.Context) {
	var records []models.EmployeeRecord

	summary := domain.GetSummary(records, c)
	if len(summary) == 0 {
		c.JSON(http.StatusNotFound, "Records not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": summary})
	}
}

func GetEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee := domain.GetEmployee(id)

	if employee.EmployeeID == 0 {
		c.JSON(http.StatusNotFound, "Employee not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": employee})
	}
}
