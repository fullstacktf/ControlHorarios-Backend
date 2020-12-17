package tests

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/gin-gonic/gin"
)

func TestHandler() *gin.Engine {
	return routes.SetupRouter("127.0.0.1:3306")
}

func ClearTestDatabase() {
	infrastructure.DB().Exec("DELETE FROM employee_records")
	infrastructure.DB().Exec("DELETE FROM sections")
	infrastructure.DB().Exec("DELETE FROM projects")
	infrastructure.DB().Exec("DELETE FROM holidays")
	infrastructure.DB().Exec("DELETE FROM employee")
	infrastructure.DB().Exec("DELETE FROM company")
	infrastructure.DB().Exec("DELETE FROM users")
}

func CreateTestUser() {
	user := models.User{UserID: 1, Username: "John", Email: "johndoe@gmail.com", Password: "foo", Rol: "employee", Status: true}
	infrastructure.DB().Debug().Create(&user)
}
