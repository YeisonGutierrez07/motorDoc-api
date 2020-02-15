package mechanic

import (
	"net/http"

	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/global"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/entities"
	"github.com/gin-gonic/gin"
)

// Get funcion para traerla  información de los mecanicos
func Get(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

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
	user := c.MustGet("user").(*entities.User)
	mechanicData := CreateMechanic{}

	err := c.ShouldBind(&mechanicData)
	if err == nil {
		mechanic, errRegister := RegisterNewMechanic(user, mechanicData)
		if errRegister != nil {
			response := global.ResponseServices(mechanicData, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(mechanic, "200", "Se he creado el usuario con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(mechanicData, "400", err.Error())
	c.JSON(400, response)
}

// MyMechanics funcion para traerla  información de los mecanicos asiciados a el taller
func MyMechanics(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	mechanic, err := GetMyMechanic(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(mechanic, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
