package appointment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// GetAppointments funcion para traerla  las rutinas que ofrece un taller
func GetAppointments(c *gin.Context) {
	id1 := c.Query("id1")
	id2 := c.Query("id2")
	date1 := c.Query("start_date")
	date2 := c.Query("end_date")

	allRoutines, err := entities.GetAppointmentsFilter(id1, id2, date1, date2)
	if err != nil {
		response := global.ResponseServices(allRoutines, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allRoutines, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// GetAppointmentsByClient funcion para traerla  las rutinas que ofrece un taller
func GetAppointmentsByClient(c *gin.Context) {
	userID := c.Param("userID")
	fhinitial := c.Query("fhinitial")
	fhend := c.Query("fhend")

	allAppointments, err := entities.GetAppointmentsByClient(userID, fhinitial, fhend)
	if err != nil {
		response := global.ResponseServices(allAppointments, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allAppointments, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// GetAppointmentsNotAvailables funcion para traerla  las citas no disponibles
func GetAppointmentsNotAvailables(c *gin.Context) {
	workshopID := c.Param("workshopID")
	routineID := c.Param("routineID")
	fhinitial := c.Query("fhinitial")
	fhend := c.Query("fhend")

	allAppointments, err := entities.GetAppointmentsNotAvailables(workshopID, routineID, fhinitial, fhend)
	if err != nil {
		response := global.ResponseServices(allAppointments, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allAppointments, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}

// GetAppointmentsMechanics funcion para traerla  las citas no disponibles
func GetAppointmentsMechanics(c *gin.Context) {
	mechanicID := c.Param("mechanicID")
	fhinitial := c.Query("fhinitial")
	fhend := c.Query("fhend")

	allAppointments, err := entities.GetAppointmentsMechanics(mechanicID, fhinitial, fhend)
	if err != nil {
		response := global.ResponseServices(allAppointments, "400", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(allAppointments, "200", "Exito")
	c.JSON(http.StatusOK, response)
	return
}
