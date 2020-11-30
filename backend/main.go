package main

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
)

var err error

func main() {
	routes.SetupRouter().Run(":8089")
}
