/**
 * Created by Yeison Gutierrez, Jorge Canchon on 09/15/2019
 */

package main

import (
	"os"

	"github.com/YeisonGutierrez07/motorDoc-api/v1/entities"
	"github.com/motorDoc-api/routes"
	"github.com/motorDoc-api/shared"
)

func main() {
	defer shared.CloseDb()

	// inicia conexion con la base de datos
	shared.Init()

	// Correr las migraciones --------------------------------------------------------------
	// Correr migracion a base de datos
	// usuarios
	shared.GetDb().AutoMigrate(&entities.User{})

	// empresa
	shared.GetDb().AutoMigrate(entities.Company{}).
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// talleres
	shared.GetDb().AutoMigrate(entities.Workshop{}).
		AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// // rutinas
	// shared.GetDb().AutoMigrate(entities.Routine{})

	// // rutinas por mecanico
	// shared.GetDb().AutoMigrate(entities.WorkshopRoutine{}).
	// 	AddForeignKey("workshop_id", "Workshops(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("routine_id", "routines(id)", "NO ACTION", "NO ACTION")

	// Mecanicos
	shared.GetDb().AutoMigrate(entities.Mechanic{}).
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION")

	// marcas
	shared.GetDb().AutoMigrate(entities.Brand{})

	// rutinas por mecanico
	shared.GetDb().AutoMigrate(entities.Vehicle{}).
		AddForeignKey("client_id", "users(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("brand_id", "brands(id)", "NO ACTION", "NO ACTION")

	// Final de las migraciones --------------------------------------------------------------

	// Inicializa las rutas
	r := routes.InitRouter()

	//Run Server

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	r.Run(":" + port)
}
