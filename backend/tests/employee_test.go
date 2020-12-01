package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestCreateEmployeeShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/employee/4").
		Body(`{
			"User":{
			   "Username": "Ana",
			   "Email": "ana@ana.ana",
			   "Password": "123Ana",
			   "Rol": "Employee"
		   },
		   "Employee": {
			   "firstName": "ana",
			   "lastName": "ana"
		   }
	   }`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetEmployeeShouldReturn200(t *testing.T) {
	user := models.User{UserID: 666, Email: "johndoe@gmail.com", Password: "foo", Rol: "employee"}
	infrastructure.DB().Create(&user)
	employee := models.Employee{EmployeeID: 666, UserID: 666, CompanyID: 2}
	infrastructure.DB().Create(&employee)

	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/employee/666").
		Expect(t).
		Status(http.StatusOK).
		End()

	infrastructure.DB().Where("employee_id = ?", 666).Delete(models.Employee{})
	infrastructure.DB().Where("user_id = ?", 666).Delete(models.User{})
}
