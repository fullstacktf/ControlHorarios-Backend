package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
)

func main() {
	routes.SetupRouter("localhost:3306").Run(":8080")
}
