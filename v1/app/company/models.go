package company

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
