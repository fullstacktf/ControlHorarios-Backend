package controllers

import (
	"log"
	"net/http"

	"github.com/fullstacktf/ControlHorarios-Backend/api/domain"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CompanyRepository struct{}

func CreateCompany(c *gin.Context) {

	var userCompany models.UserCompany
	err := c.ShouldBindWith(&userCompany, binding.JSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)

	}

	domain.CreateUser(userCompany, c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	domain.CreateCompany(userCompany, c)

}
