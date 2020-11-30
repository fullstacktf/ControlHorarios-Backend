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
		Post("/api/user/login").
		Body(`{
			"Email": "airan@gmail.com",
			"Password": "1234"
		}`).
		Expect(t).
		Body(`{
			"Rol": "employee",
			"UserID": 4
		}`).
		Status(http.StatusOK).
		End()
}
