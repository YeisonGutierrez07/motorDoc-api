package auth

import "github.com/YeisonGutierrez07/motorDoc-api/v1/entities"

// Login funcion para validar el usuario y retornar el token y objeto de usuario
func Login(loginValidator entities.DataLogin) (entities.ResponseLogin, error) {
	response, err := entities.Login(loginValidator)
	return response, err
}
