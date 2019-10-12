package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/middlewares"
	"github.com/motorDoc-api/v1/auth"
	"github.com/motorDoc-api/v1/company"
	"github.com/motorDoc-api/v1/mechanic"
	"github.com/motorDoc-api/v1/users"
	"github.com/motorDoc-api/v1/workshop"
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

		// Companies
		companiesGroup := v1.Group("companies")
		companiesGroup.GET("/", company.GetAll)
		companiesGroup.POST("/create", company.CreateNewCompany)
		companiesGroup.PUT("/changeStatus", company.ChangeStatusCompany)
		companiesGroup.DELETE("/deleteCompany/:id", company.DeleteCompany)

		companiesGroup.GET("/myCompany", company.MyCompany)
		companiesGroup.GET("/misWorkShop", company.MisWorkShop)

		// Mechanic
		workshops := v1.Group("workshop")
		workshops.GET("/", workshop.Get)

		// Mechanic
		mechanicGroup := v1.Group("mechanic")
		mechanicGroup.POST("/create", mechanic.Create)
		mechanicGroup.GET("/", mechanic.Get)
		mechanicGroup.GET("/misMechanics", mechanic.MisMechanics)

		user := v1.Group("user")
		user.GET("/", users.GetDataUser)
	}
}
