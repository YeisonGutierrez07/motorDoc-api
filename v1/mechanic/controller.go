package mechanic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/global"
	"github.com/motorDoc-api/v1/users"
)

// Get funcion para traerla  información de los mecanicos
func Get(c *gin.Context) {
	user := c.MustGet("user").(*users.User)

	companies, err := GetMechanic(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(companies, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
