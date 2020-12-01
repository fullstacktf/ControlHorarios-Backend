package tests

import (
	"net/http"
	"testing"

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
