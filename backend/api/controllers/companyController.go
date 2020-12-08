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
	var companyDto dto.CreateCompanyRequestDto
	err := c.BindJSON(&companyDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}

	DBError := domain.CreateCompany(companyDto)

	if DBError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error saving company"})
	} else {
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "Company created successfully"})
	}
}

func GetCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	company := domain.GetCompany(id)
	if company.UserID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error finding Company"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": company})
	}
}

func CreateProject(c *gin.Context) {
	var projectDto dto.ProjectDto
	c.BindJSON(&projectDto)
	if projectDto.ProjectName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}

	companyId, _ := strconv.Atoi(c.Param("id"))
	err := domain.CreateProject(companyId, projectDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating project"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
	}
}

func CreateHoliday(c *gin.Context) {
	var holiday dto.CreateHolidaysRequestDto

	err := c.BindJSON(&holiday)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
	}
	companyId, _ := strconv.Atoi(c.Param("id"))
	DBError := domain.CreateHoliday(holiday, companyId)

	if DBError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error saving holidays"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Holidays created successfully"})
	}
}

func CreateSection(c *gin.Context) {
	var section models.Sections

	err := c.ShouldBindWith(&section, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
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

func UpdateProject(c *gin.Context) {
	var projectDto dto.ProjectDto
	c.BindJSON(&projectDto)
	if projectDto.ProjectName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := domain.UpdateProjectName(id, projectDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating project"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	}
}

func UpdateSections(c *gin.Context) {
	var sectionDto dto.UpdateSectionsRequestDto
	c.BindJSON(&sectionDto)
	if sectionDto.SectionName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := domain.UpdateSectionName(id, sectionDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating project"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	}
}

func UpdateHolidays(c *gin.Context) {
	var holidaysDto dto.UpdateHolidaysRequestDto
	c.BindJSON(&holidaysDto)
	if holidaysDto.HolidaysName == "" || holidaysDto.NewDate == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := domain.UpdateHolidaysName(id, holidaysDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating holidays"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Holidays updated successfully"})
	}
}

func DeleteHolidays(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	dbErr := domain.DeleteHolidays(id)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting holidays"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Holidays deleted successfully"})
	}
}
