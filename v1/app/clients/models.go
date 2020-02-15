package clients

import "github.com/YeisonGutierrez07/motorDoc-api/v1/entities"

// Client Modelo para los clientes
type Client struct {
	entities.BaseModel
	UserID int64 `json:"user_id" db:"user_id"`
}
