package models

import (
	"log"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/infrastructure"
)

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

type Users []User

func (c *Users) Get() error {

	rows, err := infrastructure.DB.Debug().
		Model(&User{}).
		Select(`user.UserID,
						user.Username,
						user.Email,
						user.Password,
						user.JoinedDate,
						user.Rol`).
		Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = infrastructure.DB.ScanRows(rows, &user)
		if err != nil {
			log.Println("error al bindear", err)
		} else {
			log.Printf("User:%v\n", user)
		}
		*c = append([]User(*c), user)
	}
	return nil
}
