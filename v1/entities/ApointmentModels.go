package entities

// Appointment modelo de citas
type Appointment struct {
	BaseModel
	CompanyID     int64 `json:"company_id" db:"company_id"`
	MaintenanceID int64 `json:"maintenance_id" db:"maintenance_id"`
	WorkshopID    int64 `json:"workshop_id" db:"workshop_id"`
}
