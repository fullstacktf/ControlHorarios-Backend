package infrastructure

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tabler interface {
	TableName() string
}

var database *gorm.DB

func DB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	if database == nil {
		var err error
		database, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/controlhorarios?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
		if err != nil {
			log.Fatal("error al conectar a la base de datos:", err)
		}
	}
	return database
}
