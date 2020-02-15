package users

import "github.com/motorDoc-api/v1/entities"

// RegisterNewUser Funcion para registrar los nuevos usuarios
func RegisterNewUser(dataNewUser entities.NewUser) (entities.User, error) {
	user, err := entities.RegisterNewUser(dataNewUser)
	if err != nil {
		return user, err
	}
	return user, nil
}
