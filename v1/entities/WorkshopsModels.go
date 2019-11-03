package entities

import (
	"errors"

	"github.com/motorDoc-api/shared"
)

type Workshop struct {
	BaseModel
	Address   string  `json:"address" db:"address"`
	Name      string  `json:"name" db:"name"`
	Logo      string  `json:"logo" db:"logo"`
	CompanyID int64   `json:"company_id" db:"company_id"`
	UserID    int64   `json:"user_id" db:"user_id"`
	User      User    `json:"user"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

// GetWorkshop servicio para buscar y retornar la info del taller
func GetWorkshop(user *User) (Workshop, error) {
	workshop := Workshop{}
	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return workshop, errors.New("No se encontro el este usuario como taller")
	}
	return workshop, nil
}

// CreateWorkshop servicio para crear nuevos talleres
func CreateWorkshop(userCompany *User, workshop Workshop) (Workshop, error) {

	companySearch := Company{}
	if shared.GetDb().Where("user_id=?", userCompany.ID).First(&companySearch).RecordNotFound() {
		return workshop, errors.New("El usuario no tiene una empresa asociada")
	}
	workshop.CompanyID = companySearch.ID

	err := shared.GetDb().Create(&workshop).Error
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}
