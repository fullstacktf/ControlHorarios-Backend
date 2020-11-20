package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/ControlHorarios-Backend/tree/estructura_go/backend/api/models"
)

type UsersRepository struct{}

func GetUsers(c *gin.Context) {
	var users models.Users
	users.Get()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Data"})
		log.Println("Error al bindear datos", err)
		return
	}

	user.Nuevo()
	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})

}

/*func (c UsersRepository) UpdateUser() []byte {

}

func (c UsersRepository) DeleteUser() []byte {

}
*/
