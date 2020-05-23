package routines

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// Get funcion para traerla todas las rutinas
func Get(c *gin.Context) {
	allRoutines := entities.GetAllRutines()
	response := global.ResponseServices(allRoutines, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// GetByWorkshop funcion para traerla  las rutinas de los talleres
func GetByWorkshop(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	allRoutines, err := entities.GetByWorkshop(user)
	if err != nil {
		response := global.ResponseServices(allRoutines, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allRoutines, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// GetByWorkshopID funcion para traerla  las rutinas que ofrece un taller
func GetByWorkshopID(c *gin.Context) {
	workshopID, _ := strconv.ParseInt(c.Param("workshopID"), 10, 64)

	allRoutines, err := entities.GetRoutinesByWorkshopID(workshopID)
	if err != nil {
		response := global.ResponseServices(allRoutines, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allRoutines, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// GetMechanicsByRoutine funcion para traerla  las mecanicos que trabajan en esas rutinas
func GetMechanicsByRoutine(c *gin.Context) {
	routineID := c.Param("routineID")
	workshopID := c.Param("workshopID")

	allMechanics, err := entities.GetMechanicsByRoutine(routineID, workshopID)
	if err != nil {
		response := global.ResponseServices(allMechanics, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allMechanics, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// AddRoutineByWorkshop registrar de rutinas al taller
func AddRoutineByWorkshop(c *gin.Context) {
	addRoutine := []entities.AddRoutine{}
	user := c.MustGet("user").(*entities.User)

	err := c.ShouldBind(&addRoutine)
	if err == nil {
		routines, errRegister := entities.AddRoutineByWorkshop(user, addRoutine)
		if errRegister != nil {
			response := global.ResponseServices(addRoutine, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(routines, "200", "Se han agregado las rutinas con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(addRoutine, "400", err.Error())
	c.JSON(400, response)
}

// GetTreatingMechanic funcion para traerla  las rutinas que ofrece un taller
func GetTreatingMechanic(c *gin.Context) {
	workshopID := c.Param("workshopID")
	vehicleID := c.Param("vehicleID")

	allRoutines, err := entities.GetTreatingMechanic(workshopID, vehicleID)
	if err != nil {
		response := global.ResponseServices(allRoutines, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allRoutines, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}
