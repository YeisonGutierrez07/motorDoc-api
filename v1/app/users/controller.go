package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// ServiceTest funcion para hacer pruebas del server corriendo
func ServiceTest(c *gin.Context) {
	message := "Yeison nuevo deploy :D"
	response := global.ResponseServices("", "200", message)
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
