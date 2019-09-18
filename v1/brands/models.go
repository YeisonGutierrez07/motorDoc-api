package brands

import "github.com/motorDoc-api/v1/global"

// Brand modelo de marcas de vehiculos
type Brand struct {
	global.BaseModel
	Name string `json:"name" db:"name"`
}
