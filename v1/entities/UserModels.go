package entities

import (
	"errors"
	"strings"

	"github.com/motorDoc-api/shared"
)

const (
	ROLE_SUPERADMIN = "SUPERADMIN"
	ROLE_COMPANY    = "COMPANY"
	ROLE_WORKSHOP   = "WORKSHOP"
	ROLE_MECHANIC   = "MECHANIC"
	ROLE_CLIENT     = "CLIENT"
)

// User modelo de usuario
type User struct {
	BaseModel
	Name        string `json:"name" db:"name"`
	LastName    string `json:"last_name" db:"last_name"`
	Password    string `json:"-" db:"password" gorm:"not null"`
	Email       string `json:"email" db:"email" gorm:"UNIQUE_INDEX;NOT NULL"`
	Address     string `json:"address" db:"address"`
	City        string `json:"city" db:"city"`
	Role        string `json:"role" db:"role"`
	MobilePhone int64  `json:"mobilephone" db:"mobile_phone" gorm:"not null;"`
	Credential  string `json:"credential" db:"credential" gorm:"not null;UNIQUE_INDEX"`
	ProfilePic  string `json:"profile_pic" db:"profile_pic"`
}

// NewUser modelo de creación de una nueva cuenta
type NewUser struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	MobilePhone int64  `json:"mobilephone" binding:"required"`
	Credential  string `json:"credential" binding:"required"`
	ProfilePic  string `json:"profile_pic" binding:"required"`
	Role        string `json:"role" binding:"required"`
}

// DataLogin modelo que recibe el servicio del login
type DataLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// ResponseLogin modelo que responde cuando hace login el usuario
type ResponseLogin struct {
	User  User   `json:"user" db:"user"`
	Token string `json:"token"`
}

// GetUser función para retirnar un usuario que se busca por ID
func GetUser(userID int64) (User, error) {
	user := User{}
	if shared.GetDb().Set("gorm:auto_preload", true).Where("id = ?", userID).First(&user).RecordNotFound() {
		return user, errors.New("No se encontro ningun usuario")
	}
	return user, nil
}

// Login funcion para validar el usuario y retornar el token y objeto de usuario
func Login(loginValidator DataLogin) (ResponseLogin, error) {
	responseLogin := ResponseLogin{}
	passwordHashed := EncryptPassword(loginValidator.Password)
	user := User{}

	if !shared.GetDb().Where("email = ?", loginValidator.Email).First(&user).RecordNotFound() {
		if passwordHashed == user.Password {
			token, err := GenerateToken([]byte(SigningKey), user.ID, user.Credential)
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

// RegisterNewUser Funcion para registrar los nuevos usuarios
func RegisterNewUser(dataNewUser NewUser) (User, error) {
	user := User{}

	user.Name = strings.TrimSpace(dataNewUser.Name)
	user.LastName = strings.TrimSpace(dataNewUser.LastName)
	user.Password = EncryptPassword(strings.TrimSpace(dataNewUser.Password))
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
