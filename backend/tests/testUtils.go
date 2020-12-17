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

func CreateTestUser(rol string) {
	user := models.User{UserID: 1, Username: "John", Email: "johndoe@gmail.com", Password: "foo", Rol: rol, Status: true}
	infrastructure.DB().Debug().Create(&user)
}

func CreateTestEmployee() {
	CreateTestCompany()
	user := models.User{UserID: 2, Username: "Julia", Email: "juliadoe@gmail.com", Password: "foo", Rol: "employee", Status: true}
	infrastructure.DB().Debug().Create(&user)

	employee := models.Employee{EmployeeID: 1, UserID: 2, FirstName: "Julia", LastName: "Doe", CompanyID: 1}
	infrastructure.DB().Debug().Create(&employee)
}

func CreateTestCompany() {
	CreateTestUser("company")

	company := models.Company{CompanyID: 1, UserID: 1, CompanyName: "CompanyTest", Location: "Spain"}
	infrastructure.DB().Debug().Create(&company)
}

func CreateTestHolidays() {
	CreateTestCompany()
	holidays := models.Holidays{HolidayID: 1, HolidayTitle: "Navidad", HolidayDate: "2020-12-12"}
	infrastructure.DB().Debug().Create(&holidays)
}

func CreateTestSection() {
	CreateTestCompany()
	section := models.Sections{SectionID: 1, SectionName: "SectionTest", CompanyID: 1}
	infrastructure.DB().Debug().Create(&section)
}

func CreateTestProject() {
	CreateTestCompany()
	project := models.Projects{ProjectID: 1, ProjectName: "League of Legends", CompanyID: 1}
	infrastructure.DB().Create(&project)
}
