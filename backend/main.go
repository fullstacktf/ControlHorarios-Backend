package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
)

func main() {
	routes.SetupRouter("127.0.0.1:3306").Run(":8080")
}
