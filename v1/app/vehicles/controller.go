package vehicles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// GetAllVehicles funcion para traerla  información de todos los vehiculos
func GetAllVehicles(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	vehicles := entities.GetAllVehicles(user)
	response := global.ResponseServices(vehicles, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// Create registrar de nuevos vehiculos
func Create(c *gin.Context) {
	vehicleData := CreateVehicle{}
	user := c.MustGet("user").(*entities.User)

	err := c.ShouldBind(&vehicleData)
	if err == nil {
		r, errRegister := RegisterNewVehicle(user, vehicleData)
		if errRegister != nil {
			response := global.ResponseServices(vehicleData, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(r, "200", "Se he creado el vehiculo con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(vehicleData, "400", err.Error())
	c.JSON(400, response)
}
