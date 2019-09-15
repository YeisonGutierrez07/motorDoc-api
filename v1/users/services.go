package users

import (
	"strings"

	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/global"
)

// RegisterNewUser Funcion para registrar los nuevos usuarios
func RegisterNewUser(dataNewUser NewUser) (User, error) {
	user := User{}

	user.Name = strings.TrimSpace(dataNewUser.Name)
	user.LastName = strings.TrimSpace(dataNewUser.LastName)
	user.Password = global.EncryptPassword(strings.TrimSpace(dataNewUser.Password))
	user.Email = strings.TrimSpace(dataNewUser.Email)
	user.Address = strings.TrimSpace(dataNewUser.Address)
	user.City = strings.TrimSpace(dataNewUser.City)
	user.MobilePhone = strings.TrimSpace(dataNewUser.MobilePhone)
	user.Credential = strings.TrimSpace(dataNewUser.Credential)
	user.ProfilePic = strings.TrimSpace(dataNewUser.ProfilePic)
	user.Longitude = dataNewUser.Longitude
	user.Latitude = dataNewUser.Latitude

	err := shared.GetDb().Create(&user).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}