/**
 * Created by Yeison Gutierrez, Jorge Canchon on 09/15/2019
 */

package main

import (
	"github.com/motorDoc-api/routes"
	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/users"
)

func main() {
	defer shared.CloseDb()

	// inicia conexion con la base de datos
	shared.Init()

	// Correr migracion a base de datos
	shared.GetDb().AutoMigrate(&users.User{})

	// Inicializa las rutas
	r := routes.InitRouter()

	//Run Server
	r.Run(":5000")
}
