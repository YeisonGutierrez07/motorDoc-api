package mechanic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/app/users"
)

// Get funcion para traerla  información de los mecanicos
func Get(c *gin.Context) {
	user := c.MustGet("user").(*users.User)

	mechanic, err := GetMechanic(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(mechanic, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// Create registrar de nuevos mecanicos
func Create(c *gin.Context) {
	user := c.MustGet("user").(*users.User)
	mechanicData := CreateMechanic{}

	err := c.ShouldBind(&mechanicData)
	if err == nil {
		errRegister := RegisterNewMechanic(user, mechanicData)
		if errRegister != nil {
			response := global.ResponseServices(mechanicData, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(mechanicData, "200", "Se he creado el usuario con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(mechanicData, "400", err.Error())
	c.JSON(400, response)
}

// MisMechanics funcion para traerla  información de los mecanicos asiciados a el taller
func MisMechanics(c *gin.Context) {
	user := c.MustGet("user").(*users.User)

	mechanic, err := GetMisMechanic(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(mechanic, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
