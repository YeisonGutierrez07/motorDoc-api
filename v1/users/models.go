package users

import "github.com/motorDoc-api/v1/global"

// User modelo de usuario
type User struct {
	global.BaseModel
	Name        string  `json:"name" db:"name"`
	LastName    string  `json:"last_name" db:"last_name"`
	Password    string  `json:"-" db:"password" gorm:"not null"`
	Email       string  `json:"email" db:"email" gorm:"UNIQUE_INDEX;NOT NULL"`
	Address     string  `json:"address" db:"address"`
	City        string  `json:"city" db:"city"`
	MobilePhone string  `json:"mobilephone" db:"mobile_phone" gorm:"not null;"`
	Credential  string  `json:"credential" db:"credential" gorm:"not null;UNIQUE_INDEX"`
	ProfilePic  string  `json:"profile_pic" db:"profile_pic"`
	Latitude    float64 `json:"latitude" db:"latitude"`
	Longitude   float64 `json:"longitude" db:"longitude"`
}

// NewUser modelo de creación de una nueva cuenta
type NewUser struct {
	global.BaseModel
	Name        string  `json:"name" db:"name" binding:"required"`
	LastName    string  `json:"last_name" db:"last_name" binding:"required"`
	Password    string  `json:"password" db:"password" gorm:"not null" binding:"required"`
	Email       string  `json:"email" db:"email" gorm:"UNIQUE_INDEX;NOT NULL" binding:"required"`
	Address     string  `json:"address" db:"address" binding:"required"`
	City        string  `json:"city" db:"city" binding:"required"`
	MobilePhone string  `json:"mobilephone" db:"mobile_phone" gorm:"not null;" binding:"required"`
	Credential  string  `json:"credential" db:"credential" gorm:"not null;UNIQUE_INDEX" binding:"required"`
	ProfilePic  string  `json:"profile_pic" db:"profile_pic" binding:"required"`
	Latitude    float64 `json:"latitude" db:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" db:"longitude" binding:"required"`
}
