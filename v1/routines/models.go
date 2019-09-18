package routines

import "github.com/motorDoc-api/v1/global"

// Routine Modelo para los rutinas
type Routine struct {
	global.BaseModel
	Name string `json:"name" db:"name"`
}
