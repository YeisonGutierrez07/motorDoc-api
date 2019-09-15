package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/global"
)

func GetAllUsers(c *gin.Context) {
	user := User{}

	response := global.ResponseServices(user, "200", "JAAJAJJAA OK ")
	c.JSON(http.StatusOK, response)
}
