package models

import "time"

type User struct {
	UserID     int       `gorm:"column:user_id;primaryKey"`
	Username   string    `gorm:"type:varchar(50); NOT NULL"`
	Email      string    `gorm:"type:varchar(50); NOT NULL"`
	Password   string    `gorm:"type:varchar(255); NOT NULL"`
	JoinedDate time.Time `gorm:"type:date; NOT NULL"`
	Rol        string    `gorm:"type:varchar(50); NOT NULL"`
	Status     bool      `gorm:"column:status;type:boolean;"`
}

func (User) TableName() string {
	return "users"
}
