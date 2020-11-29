package tests

import (
	"net/http"
	"testing"

	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompanyShouldReturn200(t *testing.T) {
	assert.NotEqual(t, 1, 2, "No son iguales")
}

func TestCreateProjectShouldReturn200(t *testing.T) {
	apitest.New().
		Handler(routes.SetupRouter()).
		Post("/api/companies/1/projects").
		Body(`{"ProjectName": "Control Horarios 2"}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
}
