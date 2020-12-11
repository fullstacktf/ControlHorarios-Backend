package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestCreateProjectShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Post("/api/companies/1/projects").
		Body(`{"ProjectName": "Control Horarios 2"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetHolidaysShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Get("/api/companies/1/holidays").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetEmployeesShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Get("/api/companies/1/employees").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProjectsShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Get("/api/companies/1/projects").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetSectionsShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Get("/api/companies/1/sections").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateProjectNameShouldReturn200(t *testing.T) {
	project := models.Projects{ProjectID: 1, ProjectName: "League of Legends", CompanyID: 2}
	infrastructure.DB().Create(&project)

	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Put("/api/companies/1/projects").
		Body(`{"ProjectName": "Liga de Leyendas"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	infrastructure.DB().Where("project_id = ?", 1).Delete(models.Projects{})
}

func TestUpdateSectionNameShouldReturn200(t *testing.T) {
	project := models.Sections{SectionID: 666, SectionName: "Programming", CompanyID: 2}
	infrastructure.DB().Create(&project)

	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Put("/api/companies/666/sections").
		Body(`{"SectionName": "Management"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	infrastructure.DB().Where("section_id = ?", 666).Delete(models.Sections{})
}

func TestUpdateHolidayNameShouldReturn200(t *testing.T) {
	holidays := models.Holidays{HolidayID: 666, HolidayTitle: "Navidad", HolidayDate: "2022-10-12", CompanyID: 2}
	infrastructure.DB().Create(&holidays)

	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Put("/api/companies/666/holidays").
		Body(`{"HolidaysName": "Carnaval",
			  "NewDate": "2023-10-12"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	infrastructure.DB().Where("holiday_id = ?", 666).Delete(models.Holidays{})
}

func TestDeleteHolidayShouldReturn200(t *testing.T) {
	holidays := models.Holidays{HolidayID: 666, HolidayTitle: "Navidad", HolidayDate: "2022-10-12", CompanyID: 2}
	infrastructure.DB().Create(&holidays)

	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Delete("/api/companies/666/holidays").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetCompanyByIdShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Get("/api/companies/1").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateCompanyShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
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
}

func TestCreateSectionShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Post("/api/companies/1/sections").
		Body(`{"SectionName": "Testing"}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateHolidayShouldReturn201(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Post("/api/companies/1/holidays").
		Body(`{"Title": "Testing",
			   "Date": "2006-01-02"
		}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestUpdateEmployeeStatusShouldReturn200t(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Put("/api/companies/1/employees").
		Body(`{"UserID": 1}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}
