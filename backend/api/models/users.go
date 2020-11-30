package models

import (
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
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

func (c *User) Update(id int) error {
	result := infrastructure.DB.Debug().
		Model(&User{}).
		Where("users.user_id = ?", id).
		Updates(User{
			Username: c.Username,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (c *User) Delete(id int) error {
	result := infrastructure.DB.Delete(&User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
