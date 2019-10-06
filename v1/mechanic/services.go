package mechanic

import (
	"errors"

	"github.com/motorDoc-api/v1/workshop"

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

// RegisterNewMechanic Funcion para registrar las nuevos mecanicos
func RegisterNewMechanic(userWorkShop *users.User, dataNewCompany CreateMechanic) error {
	newUser := users.NewUser{}
	mechanic := Mechanic{}
	workshop := workshop.Workshop{}

	if shared.GetDb().Where("user_id = ?", userWorkShop.ID).First(&workshop).RecordNotFound() {
		return errors.New("No se encontro el este usuario como taller")
	}

	newUser.Name = dataNewCompany.Name
	newUser.LastName = dataNewCompany.LastName
	newUser.Password = dataNewCompany.Credential
	newUser.Email = dataNewCompany.Email
	newUser.Address = dataNewCompany.Address
	newUser.MobilePhone = dataNewCompany.MobilePhone
	newUser.Credential = dataNewCompany.Credential
	newUser.ProfilePic = dataNewCompany.ProfilePic
	newUser.Role = "MECHANIC"

	user, errRegister := users.RegisterNewUser(newUser)
	if errRegister != nil {
		return errRegister
	}

	mechanic.UserID = user.ID
	mechanic.CompanyID = workshop.CompanyID
	mechanic.WorkshopID = userWorkShop.ID

	err := shared.GetDb().Create(&mechanic).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMisMechanic servicio para buscar y retornar los mecanicos asociados al taller
func GetMisMechanic(user *users.User) ([]Mechanic, error) {
	workshop := workshop.Workshop{}
	mechanics := []Mechanic{}

	if shared.GetDb().Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return mechanics, errors.New("No se encontro el este usuario como taller")
	}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id = ?", workshop.ID).Find(&mechanics).RecordNotFound() {
		return mechanics, errors.New("No se encontro un listado de mecanicos")
	}

	return mechanics, nil
}
