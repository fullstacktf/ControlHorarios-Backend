package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestUserLoginShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Post("/api/user/login").
		Body(`{
			"Email": "johndoe@gmail.com",
			"Password": "foo"
		}`).
		Expect(t).
		Body(`{
			"Rol": "employee",
			"SecondaryID": 2,
			"UserID": 1
		}`).
		Status(http.StatusOK).
		End()
}

func TestUpdateUserPasswordShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter("127.0.0.1:3306")).
		Put("/api/employee/1/password").
		Body(`{
			"password":"test"
	   	}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}
