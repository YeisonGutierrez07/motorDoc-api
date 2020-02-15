package mechanic

import (
	"github.com/motorDoc-api/v1/entities"
)

// GetMechanic servicio para buscar y retornar la info del mecanico
func GetMechanic(user *entities.User) (entities.Mechanic, error) {
	mechanic, err := entities.GetMechanic(user)
	return mechanic, err
}

// RegisterNewMechanic Funcion para registrar las nuevos mecanicos
func RegisterNewMechanic(userWorkShop *entities.User, dataNewCompany CreateMechanic) (entities.Mechanic, error) {
	newUser := entities.NewUser{}
	mechanic := entities.Mechanic{}

	workshop, err := entities.GetWorkshop(userWorkShop)
	if err != nil {
		return mechanic, err
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

	user, errRegister := entities.RegisterNewUser(newUser)
	if errRegister != nil {
		return mechanic, errRegister
	}

	mechanic.UserID = user.ID
	mechanic.CompanyID = workshop.CompanyID

	mechanic, errorCreate := entities.CreateMechanic(userWorkShop, mechanic)
	return mechanic, errorCreate
}

// GetMyMechanic servicio para buscar y retornar los mecanicos asociados al taller
func GetMyMechanic(user *entities.User) ([]entities.Mechanic, error) {
	mechanics, err := entities.GetMyMechanic(user)
	return mechanics, err
}
