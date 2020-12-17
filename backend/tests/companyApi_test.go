package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestCreateProjectShouldReturn200(t *testing.T) {
	CreateTestCompany()
	apitest.New().
		Handler(TestHandler()).
		Post("/api/companies/1/projects").
		Body(`{"ProjectName": "Control Horarios 2"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetHolidaysShouldReturn200(t *testing.T) {
	CreateTestHolidays()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/companies/1/holidays").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetEmployeesShouldReturn200(t *testing.T) {
	CreateTestEmployee()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/companies/1/employees").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetProjectsShouldReturn200(t *testing.T) {
	CreateTestProject()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/companies/1/projects").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetSectionsShouldReturn200(t *testing.T) {
	CreateTestSection()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/companies/1/sections").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestUpdateProjectNameShouldReturn200(t *testing.T) {
	CreateTestProject()
	apitest.New().
		Handler(TestHandler()).
		Put("/api/companies/1/projects").
		Body(`{"ProjectName": "Liga de Leyendas"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestUpdateSectionNameShouldReturn200(t *testing.T) {
	CreateTestSection()
	apitest.New().
		Handler(TestHandler()).
		Put("/api/companies/1/sections").
		Body(`{"SectionName": "Management"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestUpdateHolidayNameShouldReturn200(t *testing.T) {
	CreateTestHolidays()
	apitest.New().
		Handler(TestHandler()).
		Put("/api/companies/1/holidays").
		Body(`{"HolidaysName": "Carnaval",
			  "NewDate": "2023-10-12"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestDeleteHolidayShouldReturn200(t *testing.T) {
	CreateTestHolidays()
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Delete("/api/companies/1/holidays").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetCompanyByIdShouldReturn200(t *testing.T) {
	CreateTestCompany()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/companies/1").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestCreateCompanyShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(TestHandler()).
		Post("/api/companies/").
		Body(`{
			   "Username": "eoi",
			   "Email": "eoi@eoi.eoi",
			   "Password": "123eoi",
			   "Rol": "Company",
			   "CompanyName": "eoi",
			   "Location": "eoi"
	   }`).
		Expect(t).
		Status(http.StatusCreated).
		End()
	ClearTestDatabase()
}

func TestCreateSectionShouldReturn201(t *testing.T) {
	CreateTestCompany()
	apitest.New().
		Handler(TestHandler()).
		Post("/api/companies/1/sections").
		Body(`{"SectionName": "Testing"}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
	ClearTestDatabase()
}

func TestCreateHolidayShouldReturn201(t *testing.T) {
	CreateTestCompany()
	apitest.New().
		Handler(TestHandler()).
		Post("/api/companies/1/holidays").
		Body(`{"Title": "Testing",
			   "Date": "2006-01-02"
		}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
	ClearTestDatabase()
}

func TestUpdateEmployeeStatusShouldReturn200(t *testing.T) {
	CreateTestEmployee()
	apitest.New().
		Handler(TestHandler()).
		Put("/api/companies/1/employees").
		Body(`{"UserID": 1,
				"Status": false}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}
