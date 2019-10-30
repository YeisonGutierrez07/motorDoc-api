package workshop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// Get funcion para traerla  información de los talleres
func Get(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	workshop, err := GetWorkshop(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(workshop, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// Create registrar de nuevos talleres
func Create(c *gin.Context) {
	workshopData := CreateWorkShop{}
	user := c.MustGet("user").(*entities.User)

	err := c.ShouldBind(&workshopData)
	if err == nil {
		r, errRegister := RegisterNewWorkShop(user, workshopData)
		if errRegister != nil {
			response := global.ResponseServices(workshopData, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(r, "200", "Se he creado el taller con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(workshopData, "400", err.Error())
	c.JSON(400, response)
}
