package workshop

import (
	"github.com/motorDoc-api/v1/global"
	"github.com/motorDoc-api/v1/users"
)

type Workshop struct {
	global.BaseModel
	Address   string     `json:"address" db:"address"`
	Name      string     `json:"name" db:"name"`
	Logo      string     `json:"logo" db:"logo"`
	CompanyID int64      `json:"company_id" db:"company_id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	User      users.User `json:"user"`
	Latitude  float64    `json:"latitude" db:"latitude"`
	Longitude float64    `json:"longitude" db:"longitude"`
}
