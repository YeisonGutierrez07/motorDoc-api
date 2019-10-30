package entities

// Vehicle Modelo para los vehiculos
type Vehicle struct {
	BaseModel
	ClientID int64  `json:"client_id" db:"client_id"`
	BrandID  int64  `json:"brand_id" db:"brand_id"`
	Model    string `json:"model" db:"model"`
}
