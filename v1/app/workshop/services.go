package workshop

import (
	"strings"

	"github.com/motorDoc-api/v1/entities"
)

// GetWorkshop servicio para buscar y retornar la info del taller
func GetWorkshop(user *entities.User) (entities.Workshop, error) {
	workshop, err := entities.GetWorkshop(user)
	return workshop, err
}

// RegisterNewWorkShop Funcion para registrar las nuevos talleres
func RegisterNewWorkShop(userCompany *entities.User, dataNewWorkshop CreateWorkShop) (entities.User, error) {
	newUser := entities.NewUser{}
	workshop := entities.Workshop{}

	newUser.Name = dataNewWorkshop.Name
	newUser.LastName = dataNewWorkshop.LastName
	newUser.Password = dataNewWorkshop.Credential
	newUser.Email = dataNewWorkshop.Email
	newUser.Address = dataNewWorkshop.Address
	newUser.MobilePhone = dataNewWorkshop.MobilePhone
	newUser.Credential = dataNewWorkshop.Credential
	newUser.ProfilePic = dataNewWorkshop.ProfilePic
	newUser.Role = "WORKSHOP"

	user, errRegister := entities.RegisterNewUser(newUser)
	if errRegister != nil {
		return user, errRegister
	}

	workshop.UserID = user.ID
	workshop.Address = strings.TrimSpace(dataNewWorkshop.AddressWorkshop)
	workshop.Logo = dataNewWorkshop.ImageWorkshop
	workshop.Name = dataNewWorkshop.Name

	_, err := entities.CreateWorkshop(userCompany, workshop)
	return user, err
}
