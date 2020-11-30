package models

import (
	"time"
)

type User struct {
	UserID     int    `gorm:"column:user_id;primaryKey"`
	Username   string `gorm:"type:varchar(50); NOT NULL" json:"username"`
	Password   string `gorm:"type:varchar(255); NOT NULL" json:"password"`
	Email      string `gorm:"type:varchar(50); NOT NULL" json:"email"`
	JoinedDate time.Time
	Rol        string `gorm:"type:varchar(50); NOT NULL" json:"rol"`
}

func (User) TableName() string {
	return "users"
}
