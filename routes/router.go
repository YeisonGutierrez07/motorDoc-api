package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/middlewares"
	"github.com/motorDoc-api/v1/auth"
	"github.com/motorDoc-api/v1/users"
)

// InitRouter inicializar del paquete de routes
func InitRouter() *gin.Engine {
	r := gin.Default()
	InitMiddleware(r)
	Routes(r)
	return r
}

// InitMiddleware inicializar Middleware para validar los endPoints
func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}

// Routes agregar los endPoints
func Routes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		//// rutas publicas
		v1.POST("/register", users.Register)
		v1.POST("/login", auth.HandleLogin)
	}
	//// rutas privadas
	v1.Use(middlewares.AuthHandler(""))
	{
		user := v1.Group("user")
		user.GET("/", users.GetDataUser)
	}
}
