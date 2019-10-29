package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
)

// HandleLogin función para hacer login
func HandleLogin(c *gin.Context) {
	loginValidator := DataLogin{}
	err := c.ShouldBind(&loginValidator)
	if err == nil {
		user, errLogin := Login(loginValidator)
		if errLogin != nil {
			response := global.ResponseServices(loginValidator, "400", errLogin.Error())
			c.JSON(http.StatusOK, response)
			return
		}
		response := global.ResponseServices(user, "200", "Se he iniciado sesión con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(loginValidator, "400", err.Error())
	c.JSON(http.StatusOK, response)
}
