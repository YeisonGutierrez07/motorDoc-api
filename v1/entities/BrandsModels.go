package entities

import (
	"errors"

	"github.com/YeisonGutierrez07/motorDoc-api/shared"
)

type Brand struct {
	BaseModel
	Name string `json:"name" db:"name"`
}

func GetAllBrands() ([]Brand, error) {
	brands := []Brand{}

	if shared.GetDb().Set("gorm:auto_preload", true).Find(&brands).Order("name ASC").RecordNotFound() {
		return brands, errors.New("No se encontraron marcas")
	}
	return brands, nil
}
