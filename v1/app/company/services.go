package company

import (
	"errors"
	"strings"

	"github.com/motorDoc-api/v1/entities"
)

// GetCompanies servicio para validar que sea un SUPERADMIN quien lo este consultando y retornar un arreglo de empresas
func GetCompanies(user *entities.User, name string) ([]entities.Company, error) {
	companies := []entities.Company{}
	if user.Role != entities.ROLE_SUPERADMIN {
		return companies, errors.New("Este usuario no tiene permisos para ver este contenido")
	}

	var filterByName = ""
	if name != "" {
		filterByName = "LOWER(business_name) LIKE '%" + name + "%'"
	}
	companies, err := entities.GetCompanies(filterByName)
	return companies, err

}

// RegisterNewCompany Funcion para registrar las nuevas empresas
func RegisterNewCompany(dataNewCompany CreateCompany) (entities.User, error) {
	newUser := entities.NewUser{}
	company := entities.Company{}

	newUser.Name = dataNewCompany.Name
	newUser.LastName = dataNewCompany.LastName
	newUser.Password = dataNewCompany.Credential
	newUser.Email = dataNewCompany.Email
	newUser.Address = dataNewCompany.Address
	newUser.MobilePhone = dataNewCompany.MobilePhone
	newUser.Credential = dataNewCompany.Credential
	newUser.ProfilePic = dataNewCompany.ProfilePic
	newUser.Role = "COMPANY"

	user, errRegister := entities.RegisterNewUser(newUser)
	if errRegister != nil {
		return user, errRegister
	}

	company.UserID = user.ID
	company.Nit = strings.TrimSpace(dataNewCompany.Nit)
	company.Logo = strings.TrimSpace(dataNewCompany.ImageCompany)
	company.BusinessName = strings.TrimSpace(dataNewCompany.BusinessName)
	company.Status = 1

	_, err := entities.CreateCompany(company)
	return user, err
}

// ChangeStatusService Funcion para cambiar el estado de las empresas
func ChangeStatusService(dataNewCompany entities.ChangeStatus) (entities.Company, error) {
	company, err := entities.ChangeStatusService(dataNewCompany)
	return company, err
}

// GetCompanyByUser servicio para buscar y retornar la info de la empresa
func GetCompanyByUser(user *entities.User) (entities.Company, error) {
	company, err := entities.GetCompanyByUser(user)
	return company, err
}

// GetMyWorkshops servicio para buscar y retornar los talleres asociados a la empresa
func GetMyWorkshops(user *entities.User) ([]entities.Workshop, error) {
	workshops, err := entities.GetMyWorkshops(user)
	return workshops, err
}
