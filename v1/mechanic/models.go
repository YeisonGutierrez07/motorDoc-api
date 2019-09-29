package mechanic

import (
	"github.com/motorDoc-api/v1/company"
	"github.com/motorDoc-api/v1/global"
	"github.com/motorDoc-api/v1/users"
	"github.com/motorDoc-api/v1/workshop"
)

// Mechanic Modelo para los mecanicos
type Mechanic struct {
	global.BaseModel
	UserID     int64             `json:"user_id" db:"user_id"`
	User       users.User        `json:"user"`
	CompanyID  int64             `json:"company_id" db:"company_id"`
	Company    company.Company   `json:"company"`
	WorkshopID int64             `json:"workshop_id" db:"workshop_id"`
	Workshop   workshop.Workshop `json:"workshop"`
}

// MechanicalRoutine modelo para guardar las rutinas por mecanico
type MechanicalRoutine struct {
	global.BaseModel
	MechanicID    int64  `json:"mechanic_id" db:"mechanic_id"`
	RoutineID     int64  `json:"routine_id" db:"routine_id"`
	EstimatedTime string `json:"estimated_time" db:"estimated_time"`
}
