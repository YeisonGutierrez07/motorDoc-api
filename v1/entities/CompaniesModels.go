package entities

import (
	"errors"

	"github.com/motorDoc-api/shared"
)

// Company Modelo para los empresas
type Company struct {
	BaseModel
	UserID       int64      `json:"user_id" db:"user_id"`
	User         User       `json:"user"`
	BusinessName string     `json:"business_name" db:"business_name"`
	Nit          string     `json:"nit" db:"nit"`
	Logo         string     `json:"logo" db:"logo"`
	Status       int        `json:"status" db:"status"`
	Workshops    []Workshop `json:"workshops"`
}

type ChangeStatus struct {
	CompanyID int64 `json:"company_id" binding:"required"`
	Status    int   `json:"status"`
}

// GetCompanies servicio para validar que sea un SUPERADMIN quien lo este consultando y retornar un arreglo de empresas
func GetCompanies(filterByName string) ([]Company, error) {
	companies := []Company{}

	if shared.GetDb().Where(filterByName).Where("status=0 OR status=1").Find(&companies).RecordNotFound() {
		return companies, errors.New("No se encontro ninguna compañia")
	}
	return companies, nil
}

// CreateCompany guardar compañia en la base de datos
func CreateCompany(company Company) (Company, error) {
	err := shared.GetDb().Create(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

// ChangeStatusService Funcion para cambiar el estado de las empresas
func ChangeStatusService(dataNewCompany ChangeStatus) (Company, error) {
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
func GetCompanyByUser(user *User) (Company, error) {
	company := Company{}
	if shared.GetDb().Set("gorm:auto_preload", true).Where("user_id = ?", user.ID).First(&company).RecordNotFound() {
		return company, errors.New("No se encontro el este usuario como mecanico")
	}
	return company, nil
}

// GetMyWorkshops servicio para buscar y retornar los talleres asociados a la empresa
func GetMyWorkshops(user *User) ([]Workshop, error) {
	company := Company{}
	workshops := []Workshop{}

	if shared.GetDb().Where("user_id = ?", user.ID).First(&company).RecordNotFound() {
		return workshops, errors.New("No se encontro el este usuario como una compañia")
	}

	if shared.GetDb().Set("gorm:auto_preload", true).Where("company_id = ?", company.ID).Find(&workshops).RecordNotFound() {
		return workshops, errors.New("No se encontro un listado de mecanicos")
	}

	return workshops, nil
}
