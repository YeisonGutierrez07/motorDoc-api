package users

import (
	"strings"

	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/app/global"
)

// RegisterNewUser Funcion para registrar los nuevos usuarios
func RegisterNewUser(dataNewUser NewUser) (User, error) {
	user := User{}

	user.Name = strings.TrimSpace(dataNewUser.Name)
	user.LastName = strings.TrimSpace(dataNewUser.LastName)
	user.Password = global.EncryptPassword(strings.TrimSpace(dataNewUser.Password))
	user.Email = strings.TrimSpace(dataNewUser.Email)
	user.Address = strings.TrimSpace(dataNewUser.Address)
	user.City = "BOGOTÁ"
	user.MobilePhone = dataNewUser.MobilePhone
	user.Credential = strings.TrimSpace(dataNewUser.Credential)
	user.ProfilePic = strings.TrimSpace(dataNewUser.ProfilePic)
	user.Role = dataNewUser.Role

	err := shared.GetDb().Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
