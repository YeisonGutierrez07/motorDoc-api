package workshop

import (
	"errors"

	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/users"
)

// GetWorkshop servicio para buscar y retornar la info del taller
func GetWorkshop(user *users.User) (Workshop, error) {
	workshop := Workshop{}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return workshop, errors.New("No se encontro el este usuario como mecanico")
	}

	return workshop, nil
}
