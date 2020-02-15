package users

import (
	"net/http"

	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/global"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/entities"
	"github.com/gin-gonic/gin"
)

// ServiceTest funcion para hacer pruebas del server corriendo
func ServiceTest(c *gin.Context) {
	response := global.ResponseServices("", "200", "Servidor corriendo con exito, tomate una cervezaAAA :v")
	c.JSON(http.StatusOK, response)
}

// Register registrar nuevos usuarios administradores
func Register(c *gin.Context) {
	newUser := entities.NewUser{}
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

// ResetPassword Cambiar la contraseña de los usuarios
func ResetPassword(c *gin.Context) {
	ResetPassword := entities.ResetPassword{}
	user := c.MustGet("user").(*entities.User)
	err := c.ShouldBind(&ResetPassword)
	if err == nil {
		user, errRegister := entities.ResetPasswordUser(user, ResetPassword)
		if errRegister != nil {
			response := global.ResponseServices(ResetPassword, "400", errRegister.Error())
			c.JSON(http.StatusAccepted, response)
			return
		}
		response := global.ResponseServices(user, "200", "Se he cambiado la contraseña con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(ResetPassword, "400", err.Error())
	c.JSON(http.StatusNotAcceptable, response)
}

// GetDataUser funcion para traer toda la información del usuario
func GetDataUser(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)
	response := global.ResponseServices(user, "200", "Información del usuario")
	c.JSON(http.StatusOK, response)
}
