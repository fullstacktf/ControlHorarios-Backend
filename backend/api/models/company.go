package models

import (
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
)

type Company struct {
	CompanyID   int `gorm:"column:company_id;primaryKey"`
	UserID      int
	User        User      `gorm:"foreignKey:UserID"`
	CompanyName string    `gorm:"column:company_name;type:varchar(50); NOT NULL json:CompanyName"`
	CreatedDate time.Time `gorm:"column:created_date;created_date"`
	Location    string    `gorm:"column:location;type:varchar(50); NOT NULL json:Location"`
}

type UserCompany struct {
	User    User    `json:"User"`
	Company Company `json:"Company"`
}

func (Company) TableName() string {
	return "company"
}

func (c *Company) NewCompany() error {

	result := infrastructure.DB.Debug().Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
