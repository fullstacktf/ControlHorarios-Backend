package models

import (
	"log"
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
)

type User struct {
	UserID     int       `gorm:"column:user_id;primaryKey"`
	Username   string    `gorm:"column:username;type:varchar(50)"`
	Email      string    `gorm:"column:email;type:varchar(50)"`
	Password   string    `gorm:"column:password;type:varchar(50)"`
	JoinedDate time.Time `gorm:"column:joined_date;joined_date"`
	Rol        string    `gorm:"column:rol;type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}

type Users []User

func (c *Users) Get() error {

	rows, err := infrastructure.DB.Debug().
		Model(&User{}).
		Select(`users.user_id,
						users.username,
						users.email,
						users.password,
						users.joined_date,
						users.rol`).
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

func (c *User) Nuevo() (error, int) {
	result := infrastructure.DB.Debug().Save(c)
	if result.Error != nil {
		return result.Error, 0
	}

	return nil, c.UserID
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
