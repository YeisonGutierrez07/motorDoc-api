/**
 * Created by Yeison Gutierrez, Jorge Canchon on 09/15/2019
 */

package main

import (
	"github.com/YeisonGutierrez07/motorDoc-api/routes"
	"github.com/YeisonGutierrez07/motorDoc-api/shared"
)

func main() {
	defer shared.CloseDb()

	// inicia conexion con la base de datos
	shared.Init()

	// Correr las migraciones --------------------------------------------------------------
	// Correr migracion a base de datos
	// usuarios
	// shared.GetDb().AutoMigrate(&entities.User{})

	// // clientes
	// shared.GetDb().AutoMigrate(entities.Client{}).
	// 	AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// // empresa
	// shared.GetDb().AutoMigrate(entities.Company{}).
	// 	AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// // talleres
	// shared.GetDb().AutoMigrate(entities.Workshop{}).
	// 	AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// // rutinas por mecanico
	// shared.GetDb().AutoMigrate(entities.WorkshopRoutine{}).
	// 	AddForeignKey("workshop_id", "Workshop(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("routine_id", "routines(id)", "NO ACTION", "NO ACTION")

	// // rutinas
	// shared.GetDb().AutoMigrate(entities.Routine{})

	// // Mecanicos
	// shared.GetDb().AutoMigrate(entities.Mechanic{}).
	// 	AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION")

	// // marcas
	// shared.GetDb().AutoMigrate(entities.Brand{})

	// // rutinas por mecanico
	// shared.GetDb().AutoMigrate(entities.Vehicle{}).
	// 	AddForeignKey("client_id", "users(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("brand_id", "brands(id)", "NO ACTION", "NO ACTION")

	// // mantenimiento
	// shared.GetDb().AutoMigrate(entities.Maintenance{}).
	// 	AddForeignKey("vehicle_id", "vehicles(id)", "NO ACTION", "NO ACTION")

	// // calificación de mantenimiento
	// shared.GetDb().AutoMigrate(entities.MaintenanceRating{}).
	// 	AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION")

	// // rutina de mantenimiento
	// shared.GetDb().AutoMigrate(entities.MaintenanceRoutine{}).
	// 	AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("routine_id", "routines(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("mechanic_id", "mechanics(id)", "NO ACTION", "NO ACTION")

	// // citas
	// shared.GetDb().AutoMigrate(entities.Appointment{}).
	// 	AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("workshop_id", "workshops(id)", "NO ACTION", "NO ACTION").
	// 	AddForeignKey("client_id", "clients(id)", "NO ACTION", "NO ACTION")
	// Final de las migraciones --------------------------------------------------------------

	// Inicializa las rutas
	r := routes.InitRouter()

	//Run Server
	r.Run(":8000")
}
