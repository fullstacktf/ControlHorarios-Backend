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
			"Email": "johndoe@gmail.com",
			"Password": "foo"
		}`).
		Expect(t).
		Body(`{
			"Rol": "employee",
			"UserID": 1
		}`).
		Status(http.StatusOK).
		End()
}
