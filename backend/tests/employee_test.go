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

func TestGetSummaryShouldReturn200(t *testing.T) {

	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/employee/2/summary").
		Expect(t).
		Status(http.StatusOK).
		End()

}

func TestDoCheckInShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/employee/2/checkin").
		Body(`{
			"description":"Test del checkin"
	   }`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDoCheckOutShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Put("/api/employee/2/checkout/70").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateEmployeePasswordShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Put("/api/employee/2/password").
		Body(`{
			"password":"test"
	   	}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}
