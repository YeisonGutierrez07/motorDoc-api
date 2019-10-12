package company

import (
	"errors"
	"strings"

	"github.com/motorDoc-api/v1/workshop"

	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/users"
)

// GetCompanies servicio para validar que sea un SUPERADMIN quien lo este consultando y retornar un arreglo de empresas
func GetCompanies(user *users.User, name string) ([]Company, error) {
	companies := []Company{}
	if user.Role != users.ROLE_SUPERADMIN {
		return companies, errors.New("Este usuario no tiene permisos para ver este contenido")
	}

	var filterByName = ""

	if name != "" {
		filterByName = "LOWER(business_name) LIKE '%" + name + "%'"
	}
	shared.GetDb().Where(filterByName).Where("status=0 OR status=1").Find(&companies)
	return companies, nil
}

// RegisterNewCompany Funcion para registrar las nuevas empresas
func RegisterNewCompany(dataNewCompany CreateCompany) (users.User, error) {
	newUser := users.NewUser{}
	company := Company{}

	newUser.Name = dataNewCompany.Name
	newUser.LastName = dataNewCompany.LastName
	newUser.Password = dataNewCompany.Credential
	newUser.Email = dataNewCompany.Email
	newUser.Address = dataNewCompany.Address
	newUser.MobilePhone = dataNewCompany.MobilePhone
	newUser.Credential = dataNewCompany.Credential
	newUser.ProfilePic = dataNewCompany.ProfilePic
	newUser.Role = "COMPANY"

	user, errRegister := users.RegisterNewUser(newUser)
	if errRegister != nil {
		return user, errRegister
	}

	company.UserID = user.ID
	company.Nit = strings.TrimSpace(dataNewCompany.Nit)
	company.Logo = strings.TrimSpace(dataNewCompany.ImageCompany)
	company.BusinessName = strings.TrimSpace(dataNewCompany.BusinessName)
	company.Status = 1

	err := shared.GetDb().Create(&company).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// ChangeStatusService Funcion para cambiar el estado de las empresas
func ChangeStatusService(dataNewCompany changeStatus) (Company, error) {
	company := Company{}
	if shared.GetDb().Where("id = ?", dataNewCompany.CompanyID).First(&company).RecordNotFound() {
		return company, errors.New("Empresa no encontrada")
	}
	company.Status = dataNewCompany.Status

	err := shared.GetDb().Save(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

// GetCompanyByUser servicio para buscar y retornar la info de la empresa
func GetCompanyByUser(user *users.User) (Company, error) {
	company := Company{}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&company).RecordNotFound() {
		return company, errors.New("No se encontro el este usuario como mecanico")
	}

	return company, nil
}

// GetMisMechanic servicio para buscar y retornar los mecanicos asociados al taller
func GetMisMechanic(user *users.User) ([]workshop.Workshop, error) {
	company := Company{}
	workshops := []workshop.Workshop{}

	if shared.GetDb().Where("user_id = ?", user.ID).First(&company).RecordNotFound() {
		return workshops, errors.New("No se encontro el este usuario como una compañia")
	}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("company_id = ?", company.ID).Find(&workshops).RecordNotFound() {
		return workshops, errors.New("No se encontro un listado de mecanicos")
	}

	return workshops, nil
}
