package shared

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Init funcion que inicia la base de datos
func Init() {
	var err error
	const (
		host     = "localhost"
		user     = "postgres"
		password = "root"
		dbname   = "biblioStore"
		database = "motorDoc"
	)

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=disable",
		user,
		password,
		host,
		database,
	)
	fmt.Println(dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Error al conectarse a la base de datos")
		panic(err)
	}
	log.Println("Base de datos conectada")
}

// GetDb funcion que trae la base de datos para consumir desde los demas paquetes
func GetDb() *gorm.DB {
	return db
}

// CloseDb funcion para cerrar la base de datos
func CloseDb() {
	db.Close()
}
