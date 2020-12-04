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
	record := models.EmployeeRecord{RecordID: 70, Description: "Testing", EmployeeID: 2}
	infrastructure.DB().Create(&record)
}

func teardown() {
	infrastructure.DB().Where("email = ?", "johndoe@gmail.com").Delete(models.User{})
	infrastructure.DB().Where("email = ?", "ana@ana.ana").Delete(models.User{})
	infrastructure.DB().Where("email = ?", "eoi@eoi.eoi").Delete(models.User{})
	infrastructure.DB().Where("first_name = ?", "ana").Delete(models.Employee{})
	infrastructure.DB().Where("first_name = ?", "eoi").Delete(models.Company{})
	infrastructure.DB().Where("project_name = ?", "Control Horarios 2").Delete(models.Projects{})
	infrastructure.DB().Where("employee_id = ?", 2).Delete(models.Employee{})
	infrastructure.DB().Where("record_id = ?", 70).Delete(models.EmployeeRecord{})
	infrastructure.DB().Where("section_name = ?", "Testing").Delete(models.Sections{})
}
