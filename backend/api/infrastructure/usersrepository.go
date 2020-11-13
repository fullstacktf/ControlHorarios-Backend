package infrastructure

import "encoding/json"

type UsersRepository struct{}

type Users struct {
	User    string `json: "user" `
	Message string `json: "message"`
}

func (c UsersRepository) GetUsers() []byte {
	users := Users{
		"user1",
		"hola",
	}
	j, _ := json.Marshal(users)
	return j
}
