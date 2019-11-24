package entities

import (
	"fmt"

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
	fmt.Println("workshopID: ", workshopID)
	fmt.Println("workshopID: ", workshopID)
	fmt.Println("workshopID: ", workshopID)
	fmt.Println("workshopID: ", workshopID)
	fmt.Println("workshopID: ", workshopID)
	allRoutines := []WorkshopRoutine{}
	shared.GetDb().Set("gorm:auto_preload", true).Where("workshop_id=?", workshopID).Find(&allRoutines)
	return allRoutines, nil
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
