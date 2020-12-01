package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestCreateProjectShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/companies/1/projects").
		Body(`{"ProjectName": "Control Horarios 2"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetHolidaysShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/companies/1/holidays").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetEmployeesShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/companies/1/employees").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProjectsShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/companies/1/projects").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetSectionsShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/companies/1/sections").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetCompanyByIdShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Get("/api/companies/1").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateCompnayShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/companies/").
		Body(`{
			"User":{
			   "Username": "eoi",
			   "Email": "eoi@eoi.eoi",
			   "Password": "123eoi",
			   "Rol": "Company"
		   },
		   "Company": {
			   "CompanyName": "eoi",
			   "Location": "eoi"
		   }
	   }`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateSectionShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/companies/1/sections").
		Body(`{"SectionName": "Testing"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateHolidayShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/companies/1/sections").
		Body(`{"holidayTitle": "Testing",
				"holidayDate": "2006-01-02T15:04:05Z"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}
