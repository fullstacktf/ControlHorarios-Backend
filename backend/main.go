package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
)

func main() {
	routes.SetupRouter("0.0.0.0:3306").Run(":8089")
}
