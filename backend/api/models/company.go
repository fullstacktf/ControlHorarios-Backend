package models

import (
	"time"
)

type Company struct {
	CompanyID   int `gorm:"column:company_id;primaryKey"`
	UserID      int
	User        User      `gorm:"foreignKey:UserID"`
	CompanyName string    `gorm:"column:company_name;type:varchar(50);"`
	CreatedDate time.Time `gorm:"column:created_date;created_date"`
	Location    string    `gorm:"column:location;type:varchar(50);"`
}

func (Company) TableName() string {
	return "company"
}
