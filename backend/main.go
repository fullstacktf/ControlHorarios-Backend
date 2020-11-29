package main

import (
	"log"
	"os"
	"time"

	s "github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)
	s.DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/controlhorarios?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	routes.SetupRouter().Run(":8089")

}
