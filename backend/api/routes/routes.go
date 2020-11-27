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

	employees := r.Group("/api/employee")
	{
		// employees.GET("/:id")  
		// employees.GET("/:id/summary")

		employees.POST("/:idCompany", controllers.CreateEmployee)
		// employees.POST("/login")					// Manuel
		// employees.POST("/:id/checkin")   // Ariane

		// employees.PUT("/:id/password")		// Airan
		// employees.PUT("/:id/checkout")   // Ariane
	}

	companies := r.Group("/api/companies")
	{
		/*companies.GET("/:id/holidays")
		companies.GET("/:id/employees")
		companies.GET("/:id/projects")
		companies.GET("/:id/sections")*/
		companies.GET("/:id", controllers.GetCompany)

		companies.POST("/", controllers.CreateCompany)
		/*companies.POST("/login")				// Jaime
		companies.POST("/:id/holidays")		// Jaime
		companies.POST("/:id/projects")		// Manuel
		companies.POST("/:id/sections") 	// Airan

		companies.PUT("/:id/projects")
		companies.PUT("/:id/sections")
		companies.PUT("/:id/holidays")

		companies.DELETE("/:id/holidays")*/
	}
	return r
}
