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
		host     = "ec2-52-73-247-67.compute-1.amazonaws.com"
		user     = "hyrjphcachsrpy"
		password = "1ef60c7ca1d1737639e27fcfd40b3b7cb65669499cdc7baefa85e76274c92043"
		database = "d8e81j1jlb7c16"
	)

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s", // sslmode=disable
		user,
		password,
		host,
		database,
	)

	fmt.Println("==>", dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Error al conectarse a la base de datos")
		panic(err)
	}
	log.Println("------Base de datos conectada-----")
}

// GetDb funcion que trae la base de datos para consumir desde los demas paquetes
func GetDb() *gorm.DB {
	return db
}

// CloseDb funcion para cerrar la base de datos
func CloseDb() {
	db.Close()
}
