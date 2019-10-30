package entities

type Brand struct {
	BaseModel
	Name string `json:"name" db:"name"`
}
