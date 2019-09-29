package mechanic

import (
	"errors"

	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/users"
)

// GetMechanic servicio para buscar y retornar la info del mecanico
func GetMechanic(user *users.User) (Mechanic, error) {
	mechanic := Mechanic{}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&mechanic).RecordNotFound() {
		return mechanic, errors.New("No se encontro el este usuario como mecanico")
	}

	return mechanic, nil
}
