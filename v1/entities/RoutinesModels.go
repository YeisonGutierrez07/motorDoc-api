package entities

import (
	"github.com/motorDoc-api/shared"
)

// Routine Modelo para los rutinas
type Routine struct {
	BaseModel
	Name string `json:"name" db:"name"`
}

// AddRoutine Modelo para agregar rutinas a los talleres
type AddRoutine struct {
	RoutineID     int64  `json:"routine_id" binding:"required"`
	EstimatedTime string `json:"estimated_time" binding:"required"`
	EstimatedCost string `json:"estimated_cost" binding:"required"`
}

type TreatingMechanic struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	LastName   string `json:"last_name" db:"last_name"`
	CompanyID  int64  `json:"company_id" db:"company_id"`
	UserID     int64  `json:"user_id" db:"user_id"`
	WorkshopID int64  `json:"workshop_id" db:"workshop_id"`
}

type MechanicRoutine struct {
	Idmechanic int64  `json:"idmechanic" db:"idmechanic"`
	Name       string `json:"name" db:"name"`
	LastName   string `json:"last_name" db:"last_name"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
}

// GetAllRutines traer todas las rutinas
func GetAllRutines() []Routine {
	allRoutines := []Routine{}
	shared.GetDb().Find(&allRoutines)
	return allRoutines
}

// GetByWorkshop traer todas las rutinas del taller
func GetByWorkshop(user *User) ([]WorkshopRoutine, error) {
	allRoutines := []WorkshopRoutine{}

	workshop, err := GetWorkshop(user)
	if err != nil {
		return allRoutines, err
	}

	shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id=?", workshop.ID).Find(&allRoutines)
	return allRoutines, nil
}

// GetRoutinesByWorkshopID traer todas las rutinas del taller filtrado por id
func GetRoutinesByWorkshopID(workshopID int64) ([]WorkshopRoutine, error) {
	allRoutines := []WorkshopRoutine{}
	shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id=?", workshopID).Find(&allRoutines)
	return allRoutines, nil
}

// GetMechanicsByRoutine traer todas las rutinas del taller filtrado por id
func GetMechanicsByRoutine(routineID string, workshopID string) ([]MechanicRoutine, error) {
	allMechanics := []MechanicRoutine{}

	query := "select * from getmechanicsbyroutine(" + routineID + "," + workshopID + ")"

	shared.GetDb().Raw(query).Find(&allMechanics)
	return allMechanics, nil
}

func AddRoutineByWorkshop(user *User, allRoutinesAdd []AddRoutine) ([]WorkshopRoutine, error) {
	allRoutines := []WorkshopRoutine{}
	workshop, err := GetWorkshop(user)
	if err != nil {
		return allRoutines, err
	}

	for _, routine := range allRoutinesAdd {
		newRoutine := WorkshopRoutine{}
		newRoutine.WorkshopID = workshop.ID
		newRoutine.RoutineID = routine.RoutineID
		newRoutine.EstimatedTime = routine.EstimatedTime
		newRoutine.EstimatedCost = routine.EstimatedCost
		err := shared.GetDb().Create(&newRoutine).Error
		if err != nil {
			return allRoutines, err
		}
	}
	shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id=?", workshop.ID).Find(&allRoutines)
	return allRoutines, nil
}

// GetTreatingMechanic traer todas las rutinas del taller filtrado por id
func GetTreatingMechanic(workshopID string, vehicleID string) ([]TreatingMechanic, error) {
	treatingMechanic := []TreatingMechanic{}

	query := "select * from gettreatingmechanic(" + workshopID + "," + vehicleID + ")"
	shared.GetDb().Raw(query).Find(&treatingMechanic)
	return treatingMechanic, nil
}
