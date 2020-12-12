package routes

import (
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(host string) *gin.Engine {
	infrastructure.SetDatabaseHost(host)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}))

	user := r.Group("/api/user")
	{
		user.POST("/login", controllers.UserLogin)
	}

	employees := r.Group("/api/employee")
	{
		employees.GET("/:id", controllers.GetEmployee)
		employees.GET("/:id/summary", controllers.GetSummary)
		employees.POST("/:id", controllers.CreateEmployee)
		employees.POST("/:id/checkin", controllers.CreateCheckIn)
		employees.PUT("/:id/password", controllers.UpdateUserPassword)
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

		companies.PUT("/:id/projects", controllers.UpdateProject)
		companies.PUT("/:id/sections", controllers.UpdateSections)
		companies.PUT("/:id/holidays", controllers.UpdateHolidays)
		companies.PUT("/:id/employees", controllers.UpdateEmployeeStatus)

		companies.DELETE("/:id/holidays", controllers.DeleteHolidays)
	}
	return r
}
