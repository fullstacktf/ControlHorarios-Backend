package routes

import (
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	var userServices models.UsersServices
	v1 := r.Group("/v1")
	{
		r.GET("/users", userServices.GetUsers)
	}
	return r
}
