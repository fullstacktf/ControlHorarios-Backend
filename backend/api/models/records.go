package models

import "time"

type EmployeeRecord struct {
	RecordID    int    `gorm:"column:record_id;primaryKey;"`
	Description string `gorm:"column:description;type:varchar(50); NOT NULL json:"description"`
	StartTime   time.Time
	EndTime     time.Time `gorm:"column:end_time;end_time"`
	EmployeeID  int
	Employee    Employee `gorm:"foreignKey:EmployeeID"`
}

func (EmployeeRecord) TableName() string {
	return "employee_records"
}
