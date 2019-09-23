package company

import "github.com/motorDoc-api/v1/global"

// Company Modelo para los empresas
type Company struct {
	global.BaseModel
	UserID       int64   `json:"user_id" db:"user_id"`
	BusinessName string  `json:"business_name" db:"business_name"`
	Nit          string  `json:"nit" db:"nit"`
	Logo         string  `json:"logo" db:"logo"`
	Status       int     `json:"status" db:"status"`
	Latitude     float64 `json:"latitude" db:"latitude" binding:"required"`
	Longitude    float64 `json:"longitude" db:"longitude" binding:"required"`
}

type companyClients struct {
	ClientID  int64 `json:"client_id" db:"client_id"`
	CompanyID int64 `json:"company_id" db:"company_id"`
}

// CreateCompany Modelo para el servicio de crear empresas
type CreateCompany struct {
	Name         string `json:"name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Credential   string `json:"credential" binding:"required"`
	ProfilePic   string `json:"profile_pic" binding:"required"`
	MobilePhone  int64  `json:"mobile_phone" binding:"required"`
	BusinessName string `json:"business_name" binding:"required"`
	ImageCompany string `json:"image_company" binding:"required"`
	Nit          string `json:"nit" binding:"required"`
}

type changeStatus struct {
	CompanyID int64 `json:"company_id" binding:"required"`
	Status    int   `json:"status"`
}
