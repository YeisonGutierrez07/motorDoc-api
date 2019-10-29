package auth

import "github.com/motorDoc-api/v1/app/users"

// DataLogin modelo que recibe el servicio del login
type DataLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// ResponseLogin modelo que responde cuando hace login el usuario
type ResponseLogin struct {
	User  users.User `json:"user" db:"user"`
	Token string     `json:"token"`
}
