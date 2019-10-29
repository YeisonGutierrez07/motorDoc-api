package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
)

// Register registrar nuevos usuarios administradores
func Register(c *gin.Context) {
	newUser := NewUser{}
	err := c.ShouldBind(&newUser)
	if err == nil {
		user, errRegister := RegisterNewUser(newUser)
		if errRegister != nil {
			response := global.ResponseServices(newUser, "400", errRegister.Error())
			c.JSON(http.StatusOK, response)
			return
		}
		response := global.ResponseServices(user, "200", "Se he creado el usuario con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(newUser, "400", err.Error())
	c.JSON(http.StatusOK, response)
}

// GetDataUser funcion para traer toda la información del usuario
func GetDataUser(c *gin.Context) {
	user := c.MustGet("user").(*User)
	response := global.ResponseServices(user, "200", "Información del usuario")
	c.JSON(http.StatusOK, response)
}
