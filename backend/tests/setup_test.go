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
}

func teardown() {
	infrastructure.DB().Where("email = ?", "johndoe@gmail.com").Delete(models.User{})
	infrastructure.DB().Where("email = ?", "ana@ana.ana").Delete(models.User{})
	infrastructure.DB().Where("first_name = ?", "ana").Delete(models.Employee{})
	infrastructure.DB().Where("project_name = ?", "Control Horarios 2").Delete(models.Projects{})
}
