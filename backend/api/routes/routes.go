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
		//v1.POST("/users", controllers.CreateUser)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}

	/*employees := r.Group("/api/employee")
	{
		employees.GET("/:id")
		employees.GET("/:id/summary")

		employees.POST("/")
		employees.POST("/login")
		employees.POST("/:id/checkin")

		employees.PUT("/:id/password")
		employees.PUT("/:id/checkout")
	}*/

	companies := r.Group("/api/companies")
	{
		/*companies.GET("/:id/holidays")
		companies.GET("/:id/employees")
		companies.GET("/:id/projects")
		companies.GET("/:id/sections")*/

		companies.POST("/", controllers.CreateCompany)
		/*companies.POST("/login")
		companies.POST("/:id/holidays")
		companies.POST("/:id/projects")
		companies.POST("/:id/sections")

		companies.PUT("/:id/projects")
		companies.PUT("/:id/sections")
		companies.PUT("/:id/holidays")

		companies.DELETE("/:id/holidays")*/
	}
	return r
}
