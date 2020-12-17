package tests

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestCreateEmployeeShouldReturn201(t *testing.T) {
	CreateTestCompany()
	apitest.New().
		Handler(TestHandler()).
		Post("/api/employee/1").
		Body(`{
			   "Username": "Ana",
			   "Email": "ana@ana.ana",
			   "Password": "123Ana",
			   "Rol": "Employee",
			   "FirstName": "ana",
			   "LastName": "ana"
	   }`).
		Expect(t).
		Status(http.StatusCreated).
		End()
	ClearTestDatabase()
}

func TestGetEmployeeShouldReturn200(t *testing.T) {
	CreateTestEmployee()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/employee/1").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestGetSummaryShouldReturn200(t *testing.T) {
	CreateTestRecords()
	apitest.New().
		Handler(TestHandler()).
		Get("/api/employee/1/summary").
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestDoCheckInShouldReturn201(t *testing.T) {
	CreateTestEmployee()
	apitest.New().
		Handler(TestHandler()).
		Post("/api/employee/1/checkin").
		Body(`{
			"description":"Test del checkin"
	   }`).
		Expect(t).
		Status(http.StatusCreated).
		End()
	ClearTestDatabase()
}

func TestDoCheckOutShouldReturn200(t *testing.T) {
	CreateTestRecords()
	apitest.New().
		Handler(TestHandler()).
		Put("/api/employee/1/checkout/1").
		Expect(t).
		Status(http.StatusOK).
		End()

	ClearTestDatabase()
}
