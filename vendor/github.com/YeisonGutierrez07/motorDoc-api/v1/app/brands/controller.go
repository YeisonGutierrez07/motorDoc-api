package brands

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// GetAll funcion para traerla  información de todos las marcas
func GetAll(c *gin.Context) {
	brands, err := entities.GetAllBrands()
	if err != nil {
		response := global.ResponseServices(brands, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(brands, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
