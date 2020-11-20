package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/models"
)

type UsersRepository struct{}

func GetUsers(c *gin.Context) {
	var users models.Users
	users.Get()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser() {
	user := &models.User{
		UserID:     1,
		Username:   "Pepe",
		Email:      "pepe@pepe.pepe",
		Password:   "123pepe",
		JoinedDate: time.Now(),
		Rol:        "Employee",
	}
	infrastructure.DB.Create(user)

}

/*func (c UsersRepository) UpdateUser() []byte {

}

func (c UsersRepository) DeleteUser() []byte {

}
*/
