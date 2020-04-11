package entities

import (
	"github.com/motorDoc-api/shared"
)

// Vehicle Modelo para los vehiculos
type Vehicle struct {
	BaseModel
	ClientID  int64  `json:"client_id" db:"client_id"`
	BrandID   int64  `json:"brand_id" db:"brand_id"`
	Brand     Brand  `json:"brand"`
	Model     string `json:"model" db:"model"`
	Estado    string `json:"estado" db:"estado"`
	Image     string `json:"image" db:"image"`
	Reference string `json:"reference" db:"reference"`
	Placa     string `json:"placa" db:"placa"`
}

func GetAllVehicles(user *User) []Vehicle {
	Vehicles := []Vehicle{}
	shared.GetDb().Set("gorm:auto_preload", true).Where("client_id =?", user.ID).Find(&Vehicles)
	return Vehicles
}

func CreateVehicle(newVehicle Vehicle) (Vehicle, error) {

	err := shared.GetDb().Create(&newVehicle).Error
	if err != nil {
		return newVehicle, err
	}
	return newVehicle, nil
}
