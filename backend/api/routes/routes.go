package routes

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.CreateUser)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}
	return r
}
