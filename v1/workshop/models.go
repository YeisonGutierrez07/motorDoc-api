package workshop

import "github.com/motorDoc-api/v1/global"

type Workshop struct {
	global.BaseModel
	CompanyID int64   `json:"company_id" db:"company_id"`
	Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}
