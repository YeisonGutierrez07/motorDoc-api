/**
 * Created by Yeison Gutierrez, Jorge Canchon on 09/15/2019
 */

package main

import (
	"github.com/motorDoc-api/routes"
	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/appointment"
	"github.com/motorDoc-api/v1/brands"
	"github.com/motorDoc-api/v1/clients"
	"github.com/motorDoc-api/v1/company"
	"github.com/motorDoc-api/v1/maintenance"
	"github.com/motorDoc-api/v1/mechanic"
	"github.com/motorDoc-api/v1/routines"
	"github.com/motorDoc-api/v1/users"
	"github.com/motorDoc-api/v1/vehicles"
	"github.com/motorDoc-api/v1/workshop"
)

func main() {
	defer shared.CloseDb()

	// inicia conexion con la base de datos
	shared.Init()

	// Correr las migraciones

	// Correr migracion a base de datos
	// usuarios
	shared.GetDb().AutoMigrate(&users.User{})

	// clientes
	shared.GetDb().AutoMigrate(clients.Client{}).
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// empresa
	shared.GetDb().AutoMigrate(company.Company{}).
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION")

	// talleres
	shared.GetDb().AutoMigrate(workshop.Workshop{}).
		AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION")

	// rutinas
	shared.GetDb().AutoMigrate(routines.Routine{})

	// Mecanicos
	shared.GetDb().AutoMigrate(mechanic.Mechanic{}).
		AddForeignKey("user_id", "users(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION")

	// rutinas por mecanico
	shared.GetDb().AutoMigrate(mechanic.MechanicalRoutine{}).
		AddForeignKey("mechanic_id", "mechanics(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("routine_id", "routines(id)", "NO ACTION", "NO ACTION")

	// marcas
	shared.GetDb().AutoMigrate(brands.Brand{})

	// rutinas por mecanico
	shared.GetDb().AutoMigrate(vehicles.Vehicle{}).
		AddForeignKey("client_id", "clients(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("brand_id", "brands(id)", "NO ACTION", "NO ACTION")

	// mantenimiento
	shared.GetDb().AutoMigrate(maintenance.Maintenance{}).
		AddForeignKey("vehicle_id", "vehicles(id)", "NO ACTION", "NO ACTION")

	// calificación de mantenimiento
	shared.GetDb().AutoMigrate(maintenance.MaintenanceRating{}).
		AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION")

	// rutina de mantenimiento
	shared.GetDb().AutoMigrate(maintenance.MaintenanceRoutine{}).
		AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("routine_id", "routines(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("mechanic_id", "mechanics(id)", "NO ACTION", "NO ACTION")

	// citas
	shared.GetDb().AutoMigrate(appointment.Appointment{}).
		AddForeignKey("maintenance_id", "maintenances(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("company_id", "companies(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("workshop_id", "workshops(id)", "NO ACTION", "NO ACTION").
		AddForeignKey("client_id", "clients(id)", "NO ACTION", "NO ACTION")

	// Inicializa las rutas
	r := routes.InitRouter()

	//Run Server
	r.Run(":5000")
}
