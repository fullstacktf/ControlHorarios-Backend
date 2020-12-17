package tests

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestUserLoginShouldReturn200(t *testing.T) {
	CreateTestUser("employee")
	apitest.New().
		Handler(TestHandler()).
		Post("/api/user/login").
		Body(`{
			"Email": "johndoe@gmail.com",
			"Password": "foo"
		}`).
		Expect(t).
		Body(`{
			"Rol": "employee",
			"SecondaryID": 0,
			"UserID": 1
		}`).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}

func TestUpdateUserPasswordShouldReturn200(t *testing.T) {
	CreateTestUser("employee")
	apitest.New().
		Handler(TestHandler()).
		Put("/api/employee/1/password").
		Body(`{
			"password":"test"
	   	}`).
		Expect(t).
		Status(http.StatusOK).
		End()
	ClearTestDatabase()
}
