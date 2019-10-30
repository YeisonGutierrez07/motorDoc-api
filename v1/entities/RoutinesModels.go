package entities

// Routine Modelo para los rutinas
type Routine struct {
	BaseModel
	Name string `json:"name" db:"name"`
}
