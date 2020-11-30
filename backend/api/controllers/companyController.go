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

func CreateCompany(c *gin.Context) {
	var userCompany models.UserCompany
	err := c.ShouldBindWith(&userCompany, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	_, id := domain.CreateUser(userCompany, c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	domain.CreateCompany(userCompany, c, id)
}

func GetCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	company := domain.GetCompany(id)
	c.JSON(http.StatusOK, gin.H{"data": company})
}

func CreateProject(c *gin.Context) {
	var projectDto dto.ProjectDto
	c.BindJSON(&projectDto)
	if projectDto.ProjectName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := domain.CreateProject(id, projectDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating project"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
	}
}

func CreateHoliday(c *gin.Context) {
	var holidayCompany models.Holidays

	err := c.ShouldBindWith(&holidayCompany, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	domain.CreateHoliday(holidayCompany, c)
}

func CreateSection(c *gin.Context) {
	var section models.Sections

	err := c.ShouldBindWith(&section, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
	}

	domain.CreateSection(section, c)
}

func GetHolidays(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	holidays := domain.GetHolidays(id)
	if len(holidays) == 0 {
		c.JSON(http.StatusNotFound, "Holidays list not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": holidays})
	}
}

func GetEmployees(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	employees := domain.GetEmployees(id)
	if len(employees) == 0 {
		c.JSON(http.StatusNotFound, "Employee list not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": employees})
	}
}

func GetProjects(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	projects := domain.GetProjects(id)
	if len(projects) == 0 {
		c.JSON(http.StatusNotFound, "Project list not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": projects})
	}
}

func GetSections(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	sections := domain.GetSections(id)
	if len(sections) == 0 {
		c.JSON(http.StatusNotFound, "Project list not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sections})
	}
}
