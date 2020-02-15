package routes

import (
	"github.com/YeisonGutierrez07/motorDoc-api/middlewares"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/auth"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/brands"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/company"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/mechanic"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/routines"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/users"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/vehicles"
	"github.com/YeisonGutierrez07/motorDoc-api/v1/app/workshop"
	"github.com/gin-gonic/gin"
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

	r.GET("/", users.ServiceTest)

	v1 := r.Group("/v1")
	{
		//// rutas publicas
		v1.POST("/testv1", users.ServiceTest)
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
		companiesGroup.GET("/myWorkShop", company.MyWorkShop)

		// // Mechanic
		workshops := v1.Group("workshop")
		workshops.GET("/", workshop.Get)
		workshops.GET("/search/:workshopID", workshop.GetByID)
		workshops.GET("/all", workshop.GetAll)
		workshops.POST("/create", workshop.Create)

		// Mechanic
		mechanicGroup := v1.Group("mechanic")
		mechanicGroup.POST("/create", mechanic.Create)
		mechanicGroup.GET("/", mechanic.Get)
		mechanicGroup.GET("/myMechanics", mechanic.MyMechanics)

		// routines
		routine := v1.Group("routines")
		routine.GET("", routines.Get)
		routine.GET("/byWorkshop", routines.GetByWorkshop)
		routine.GET("/byWorkshopID/:workshopID", routines.GetByWorkshopID)

		routine.POST("/addRoutineByWorkshop", routines.AddRoutineByWorkshop)

		// users
		user := v1.Group("user")
		user.GET("/", users.GetDataUser)
		user.PUT("/reset", users.ResetPassword)

		// brands
		brand := v1.Group("brands")
		brand.GET("", brands.GetAll)

		// brands
		vehicle := v1.Group("vehicles")
		vehicle.GET("", vehicles.GetAllVehicles)
		vehicle.POST("", vehicles.Create)

	}
}
