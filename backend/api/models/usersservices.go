package models

import (
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/infrastructure"
	"github.com/gin-gonic/gin"
)

type UsersServices struct {
}

func (c UsersServices) GetUsers(context *gin.Context) {

	var usersRepository infrastructure.UsersRepository

	context.Data(200, "application/json", usersRepository.GetUsers())
}
