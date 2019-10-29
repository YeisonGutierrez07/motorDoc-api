package mechanic

import (
	"github.com/motorDoc-api/v1/app/company"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/app/users"
	"github.com/motorDoc-api/v1/app/workshop"
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

// CreateMechanic Modelo para el servicio de crear mecanicos
type CreateMechanic struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Credential  string `json:"credential" binding:"required"`
	ProfilePic  string `json:"profile_pic" binding:"required"`
	MobilePhone int64  `json:"mobile_phone" binding:"required"`
}
