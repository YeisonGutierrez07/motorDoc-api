package vehicles

import "github.com/motorDoc-api/v1/app/global"

// Vehicle Modelo para los vehiculos
type Vehicle struct {
	global.BaseModel
	ClientID int64  `json:"client_id" db:"client_id"`
	BrandID  int64  `json:"brand_id" db:"brand_id"`
	Model    string `json:"model" db:"model"`
}
