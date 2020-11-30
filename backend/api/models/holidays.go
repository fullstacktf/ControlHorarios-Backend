package models

import (
	"time"
)

type Holidays struct {
	HolidayID    int       `gorm:"column:holiday_id;primaryKey"`
	HolidayTitle string    `gorm:"column:holiday_title;type:varchar(50); NOT NULL json:"holidayTitle"`
	HolidayDate  time.Time `gorm:"column:holiday_date;holiday_date; NOT NULL json:"holidayDate"`
	CompanyID    int
	Company      Company `gorm:"foreignKey:CompanyID"`
}

func (Holidays) TableName() string {
	return "holidays"
}
