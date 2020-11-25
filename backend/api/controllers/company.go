package controllers

import (
	"log"
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

type CompanyRepository struct{}

func CreateCompany(c *gin.Context) {

	id := CreateUser(c)

	var company models.Company
	err := c.ShouldBindJSON(&company)
	company.User.ID = id
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	company.NewCompany()
	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})

}
