package infrastructure

import "gorm.io/gorm"

var DB *gorm.DB

type Tabler interface {
	TableName() string
}
