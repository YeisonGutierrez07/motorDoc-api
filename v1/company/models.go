package company

import "github.com/motorDoc-api/v1/global"

// Company Modelo para los empresas
type Company struct {
	global.BaseModel
	UserID int64 `json:"user_id" db:"user_id"`
}

type companyClients struct {
	ClientID  int64 `json:"client_id" db:"client_id"`
	CompanyID int64 `json:"company_id" db:"company_id"`
}
