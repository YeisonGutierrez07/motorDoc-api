package clients

import "github.com/motorDoc-api/v1/global"

// Client Modelo para los clientes
type Client struct {
	global.BaseModel
	UserID int64 `json:"user_id" db:"user_id"`
}