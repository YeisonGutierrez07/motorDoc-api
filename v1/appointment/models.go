package appointment

import "github.com/motorDoc-api/v1/global"

// Appointment modelo de citas
type Appointment struct {
	global.BaseModel
	CompanyID     int64 `json:"company_id" db:"company_id"`
	MaintenanceID int64 `json:"maintenance_id" db:"maintenance_id"`
	WorkshopID    int64 `json:"workshop_id" db:"workshop_id"`
	ClientID      int64 `json:"client_id" db:"client_id"`
}
