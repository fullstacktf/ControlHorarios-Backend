package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/routes"
)

func main() {
	routes.SetupRouter().Run(":8989")
}
