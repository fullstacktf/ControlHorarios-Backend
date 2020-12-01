package tests

import (
	"os"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	user := models.User{UserID: 1, Email: "johndoe@gmail.com", Password: "foo", Rol: "employee"}
	infrastructure.DB().Create(&user)
	employee := models.Employee{EmployeeID: 2, UserID: user.UserID, CompanyID: 1}
	infrastructure.DB().Create(&employee)
}

func teardown() {
	infrastructure.DB().Where("email = ?", "johndoe@gmail.com").Delete(models.User{})
	infrastructure.DB().Where("email = ?", "ana@ana.ana").Delete(models.User{})
	infrastructure.DB().Where("first_name = ?", "ana").Delete(models.Employee{})
	infrastructure.DB().Where("project_name = ?", "Control Horarios 2").Delete(models.Projects{})
	infrastructure.DB().Where("employee_id = ?", 2).Delete(models.Employee{})
}
