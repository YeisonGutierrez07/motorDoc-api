package maintenance

import (
	"time"

	"github.com/motorDoc-api/v1/global"
)

// Maintenance Modelo para los mantenimientos
type Maintenance struct {
	global.BaseModel
	VehicleID int64  `json:"vehicle_id" db:"vehicle_id"`
	TotalCost string `json:"total_cost" db:"total_cost"`
}

// MaintenanceRoutine modelo para la tabla de rutinas de mantenimiento
type MaintenanceRoutine struct {
	global.BaseModel
	MaintenanceID int64     `json:"maintenance_id" db:"maintenance_id"`
	RoutineID     int64     `json:"routine_id" db:"routine_id"`
	MechanicID    int64     `json:"mechanic_id" db:"mechanic_id"`
	EstimatedTime string    `json:"estimated_time" db:"estimated_time"`
	TotalTime     string    `json:"total_time" db:"total_time"`
	Cost          int64     `json:"cost" db:"cost"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndSate       time.Time `json:"end_date" db:"end_date"`
}

// MaintenanceRating modelo para la tabla de calificación de mantenimiento
type MaintenanceRating struct {
	global.BaseModel
	MaintenanceID int64  `json:"maintenance_id" db:"maintenance_id"`
	Stars         int    `json:"stars" db:"stars"`
	Commentary    string `json:"commentary" db:"commentary"`
}
