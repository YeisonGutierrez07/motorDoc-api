package workshop

import (
	"errors"
	"strings"

	"github.com/motorDoc-api/shared"
	// "github.com/motorDoc-api/v1/app/company"
	"github.com/motorDoc-api/v1/app/users"
)

// GetWorkshop servicio para buscar y retornar la info del taller
func GetWorkshop(user *users.User) (Workshop, error) {
	workshop := Workshop{}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&workshop).RecordNotFound() {
		return workshop, errors.New("No se encontro el este usuario como mecanico")
	}

	return workshop, nil
}

// RegisterNewWorkShop Funcion para registrar las nuevos talleres
func RegisterNewWorkShop(userCompany *users.User, dataNewWorkshop CreateWorkShop) (users.User, error) {
	newUser := users.NewUser{}
	workshop := Workshop{}

	newUser.Name = dataNewWorkshop.Name
	newUser.LastName = dataNewWorkshop.LastName
	newUser.Password = dataNewWorkshop.Credential
	newUser.Email = dataNewWorkshop.Email
	newUser.Address = dataNewWorkshop.Address
	newUser.MobilePhone = dataNewWorkshop.MobilePhone
	newUser.Credential = dataNewWorkshop.Credential
	newUser.ProfilePic = dataNewWorkshop.ProfilePic
	newUser.Role = "WORKSHOP"

	user, errRegister := users.RegisterNewUser(newUser)
	if errRegister != nil {
		return user, errRegister
	}

	// companySearch := company.Company{}
	// shared.GetDb().Where("user_id=?", userCompany.ID).First(&company)

	workshop.UserID = user.ID
	workshop.Address = strings.TrimSpace(dataNewWorkshop.AddressWorkshop)
	workshop.Logo = dataNewWorkshop.ImageWorkshop
	workshop.Name = dataNewWorkshop.Name
	// workshop.CompanyID = companySearch.ID

	err := shared.GetDb().Create(&workshop).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
