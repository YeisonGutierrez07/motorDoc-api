package entities

import (
	"errors"

	"github.com/YeisonGutierrez07/motorDoc-api/shared"
)

// Mechanic Modelo para los mecanicos
type Mechanic struct {
	BaseModel
	UserID     int64    `json:"user_id" db:"user_id"`
	User       User     `json:"user"`
	CompanyID  int64    `json:"company_id" db:"company_id"`
	Company    Company  `json:"company"`
	WorkshopID int64    `json:"workshop_id" db:"workshop_id"`
	Workshop   Workshop `json:"workshop"`
}

// GetMechanic servicio para buscar en la base de datos y retornar la info del mecanico
func GetMechanic(user *User) (Mechanic, error) {
	mechanic := Mechanic{}
	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&mechanic).RecordNotFound() {
		return mechanic, errors.New("No se encontro el este usuario como mecanico")
	}
	return mechanic, nil
}

// CreateMechanic servicio para crear en la base de datos los nuevos talleres
func CreateMechanic(user *User, mechanic Mechanic) (Mechanic, error) {
	workshop := Workshop{}

	if shared.GetDb().Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return mechanic, errors.New("No se encontro el este usuario como taller")
	}
	mechanic.WorkshopID = workshop.ID
	err := shared.GetDb().Create(&mechanic).Error
	if err != nil {
		return mechanic, err
	}
	return mechanic, nil
}

// GetMyMechanic servicio para buscar y retornar los mecanicos asociados al taller
func GetMyMechanic(user *User) ([]Mechanic, error) {
	workshop := Workshop{}
	mechanics := []Mechanic{}

	if shared.GetDb().Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return mechanics, errors.New("No se encontro el este usuario como taller")
	}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id = ?", workshop.ID).Find(&mechanics).RecordNotFound() {
		return mechanics, errors.New("No se encontro un listado de mecanicos")
	}
	return mechanics, nil
}
