package users

import "github.com/motorDoc-api/v1/global"

type User struct {
	global.BaseModel
	Name        string `json:"name" db:"name"`
	Password    string `json:"-" db:"password" gorm:"not null"`
	Email       string `json:"email" db:"email" gorm:"UNIQUE_INDEX;NOT NULL"`
	Address     string `json:"address" db:"address"`
	City        string `json:"city" db:"city"`
	MobilePhone string `json:"mobilephone" db:"mobile_phone" gorm:"not null;"`
	Credential  string `json:"credential" db:"credential" gorm:"not null;UNIQUE_INDEX"`
	ProfilePic  string `json:"profile_pic" db:"profile_pic"`
}
