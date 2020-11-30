package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
)

func TestUserLoginShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/employee/login").
		Body(`{
			"Email": "airan@gmail.com",
			"Password": "12345"
		}`).
		Expect(t).
		Body(`{
			"Rol": "employee",
			"UserID": 4
		}`).
		Status(http.StatusOK).
		End()
}
