package models

import "time"

type User struct {
	UserID     int       `gorm:"primaryKey"`
	Username   string    `gorm:"type:varchar(50)"`
	Email      string    `gorm:"type:varchar(50)"`
	Password   string    `gorm:"type:varchar(50)"`
	JoinedDate time.Time `gorm:"joined_date"`
	Rol        string    `gorm:"type:varchar(50)"`
}

func (a *User) TableName() string {
	return "User"
}
