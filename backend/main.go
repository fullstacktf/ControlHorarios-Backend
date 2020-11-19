package main

import (
	"log"
	"os"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/config"
	s "github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/models"
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/routes"
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
	s.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	s.DB.AutoMigrate(&models.TipoCoche{}, &models.Coche{}, &models.Cliente{}, &models.Alquiler{})

	routes.SetupRouter().Run(":8080")
}
