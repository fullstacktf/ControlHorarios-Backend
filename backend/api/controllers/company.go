package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

type CompanyRepository struct{}

func CreateCompany(c *gin.Context) {

	id := CreateUser(c)

	fmt.Println(id)
	var company models.Company
	err := c.BindJSON(&company)
	company.User.UserID = id
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	company.NewCompany()
	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})

}
