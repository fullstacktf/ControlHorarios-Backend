package models

import "time"

type Holidays struct {
	HolidayID    int       `gorm:"column:holiday_id;primaryKey"`
	HolidayTitle string    `gorm:"column:holiday_title;type:varchar(50)"`
	HolidayDate  time.Time `gorm:"column:holiday_date;holiday_date"`
	Company      Company   `gorm:"foreignKey:CompanyID"`
}

func (Holidays) TableName() string {
	return "holidays"
}
