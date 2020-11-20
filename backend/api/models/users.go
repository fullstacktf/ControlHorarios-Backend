package models

import (
	"log"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/infrastructure"
)

type User struct {
	UserID     int       `gorm:"column:user_id;primaryKey"`
	Username   string    `gorm:"column:username;type:varchar(50)"`
	Email      string    `gorm:"column:email;type:varchar(50)"`
	Password   string    `gorm:"column:password;type:varchar(50)"`
	JoinedDate time.Time `gorm:"column:joined_date;joined_date"`
	Rol        string    `gorm:"column:rol;type:varchar(50)"`
}

func (a *User) TableName() string {
	return "User"
}

type Users []User

func (c *Users) Get() error {

	rows, err := infrastructure.DB.Debug().
		Model(&User{}).
		Select(`User.user_id,
						User.username,
						User.email,
						User.password,
						User.joined_date,
						User.rol`).
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

func (c *User) Nuevo() error {
	result := infrastructure.DB.Debug().Save(c)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
