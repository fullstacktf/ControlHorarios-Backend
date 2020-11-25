package domain

import (
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func CreateCompany(company models.UserCompany, c *gin.Context) error {

	result := infrastructure.DB.Debug().
		Select(`Company.location, Company.company_name, User.user_id`).Create(&company.Company)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New company created successfully"})
	return nil
}
