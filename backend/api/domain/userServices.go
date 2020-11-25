package domain

import (

	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/model"
)

func Get() []User {
	type Users []model.User

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
		Users = append([]User(Users), user)
	}
	return Users
}