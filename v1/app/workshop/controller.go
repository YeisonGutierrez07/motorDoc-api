package workshop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/global"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/entities"
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

// GetByID funcion para buscar un taller por ID
func GetByID(c *gin.Context) {
	workshopID, _ := strconv.ParseInt(c.Param("workshopID"), 10, 64)

	workshop, err := entities.GetWorkshopByID(workshopID)
	if err != nil {
		response := global.ResponseServices(workshop, "400", err.Error())
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

// GetAll funcion para traerla  información de todos los talleres
func GetAll(c *gin.Context) {
	search := c.Query("search")

	workshops, err := entities.GetAllWorkshop(search)
	if err != nil {
		response := global.ResponseServices(workshops, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(workshops, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
