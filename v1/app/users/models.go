package users

import "github.com/motorDoc-api/v1/app/global"

const (
	ROLE_SUPERADMIN = "SUPERADMIN"
	ROLE_COMPANY    = "COMPANY"
	ROLE_WORKSHOP   = "WORKSHOP"
	ROLE_MECHANIC   = "MECHANIC"
	ROLE_CLIENT     = "CLIENT"
)

// User modelo de usuario
type User struct {
	global.BaseModel
	Name        string `json:"name" db:"name"`
	LastName    string `json:"last_name" db:"last_name"`
	Password    string `json:"-" db:"password" gorm:"not null"`
	Email       string `json:"email" db:"email" gorm:"UNIQUE_INDEX;NOT NULL"`
	Address     string `json:"address" db:"address"`
	City        string `json:"city" db:"city"`
	Role        string `json:"role" db:"role"`
	MobilePhone int64  `json:"mobilephone" db:"mobile_phone" gorm:"not null;"`
	Credential  string `json:"credential" db:"credential" gorm:"not null;UNIQUE_INDEX"`
	ProfilePic  string `json:"profile_pic" db:"profile_pic"`
}

// NewUser modelo de creación de una nueva cuenta
type NewUser struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	MobilePhone int64  `json:"mobilephone" binding:"required"`
	Credential  string `json:"credential" binding:"required"`
	ProfilePic  string `json:"profile_pic" binding:"required"`
	Role        string `json:"role" binding:"required"`
}
