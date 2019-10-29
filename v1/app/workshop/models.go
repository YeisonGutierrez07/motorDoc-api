package workshop

import (
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/app/users"
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

// CreateWorkShop Modelo para el servicio de crear talleres
type CreateWorkShop struct {
	Name            string `json:"name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Address         string `json:"address" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Credential      string `json:"credential" binding:"required"`
	ProfilePic      string `json:"profile_pic" binding:"required"`
	MobilePhone     int64  `json:"mobile_phone" binding:"required"`
	AddressWorkshop string `json:"address_workshop" binding:"required"`
	ImageWorkshop   string `json:"imageWorkshop" binding:"required"`
	NameWorkshop    string `json:"name_workshop" binding:"required"`
}
