package auth

import (
	"errors"

	"github.com/motorDoc-api/middlewares"
	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/global"
	"github.com/motorDoc-api/v1/users"
)

// Login funcion para validar el usuario y retornar el token y objeto de usuario
func Login(loginValidator DataLogin) (ResponseLogin, error) {
	responseLogin := ResponseLogin{}
	passwordHashed := global.EncryptPassword(loginValidator.Password)
	user := users.User{}

	if !shared.GetDb().Where("email = ?", loginValidator.Email).First(&user).RecordNotFound() {
		if passwordHashed == user.Password {
			token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), user.ID, user.Credential)
			if err == nil {
				responseLogin.User = user
				responseLogin.Token = token
				return responseLogin, nil
			}
		}
		return responseLogin, errors.New("La contraseña es invalida")
	}
	return responseLogin, errors.New("El correo no existe en la base de datos")
}
