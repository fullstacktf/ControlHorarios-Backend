package domain

import (
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func CreateCompany(company models.UserCompany, c *gin.Context, id int) error {

	company.Company.UserID = id

	result := infrastructure.DB().Debug().
		Select(`Company.location, Company.company_name`).Create(&company.Company)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New company created successfully"})
	return nil
}

func GetCompany(id int) models.Company {
	var company models.Company
	infrastructure.DB().First(&company, id)

	return company
}

func CreateProject(id int, projectDto dto.ProjectDto) error {
	return infrastructure.InsertProject(id, projectDto.ProjectName)
}
