package routes

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/api/user")
	{

		user.POST("/login", controllers.UserLogin)
	}

	employees := r.Group("/api/employee")
	{
		// employees.GET("/:id")

		// employees.GET("/:id/summary")

		employees.POST("/:id", controllers.CreateEmployee)
		employees.POST("/:id/checkin", controllers.CreateCheckIn)

		employees.PUT("/:id/password", controllers.UpdateEmployeePassword)
		employees.PUT("/:id/checkout/:idRecord", controllers.DoCheckOut)
	}

	companies := r.Group("/api/companies")
	{
		companies.GET("/:id/holidays", controllers.GetHolidays)
		companies.GET("/:id/employees", controllers.GetEmployees)
		companies.GET("/:id/projects", controllers.GetProjects)
		companies.GET("/:id/sections", controllers.GetSections)
		companies.GET("/:id", controllers.GetCompany)

		companies.POST("/", controllers.CreateCompany)

		companies.POST("/:id/projects", controllers.CreateProject)
		companies.POST("/:id/holidays", controllers.CreateHoliday)
		companies.POST("/:id/sections", controllers.CreateSection)

		/*companies.PUT("/:id/projects")
		companies.PUT("/:id/sections")
		companies.PUT("/:id/holidays")

		companies.DELETE("/:id/holidays")*/
	}
	return r
}
