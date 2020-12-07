package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
)

func main() {
	routes.SetupRouter("172.28.1.1:3306").Run(":8080")
}
