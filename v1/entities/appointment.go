package entities

import "github.com/motorDoc-api/shared"

type AppointmentsResponse struct {
	Idmechanic      int64  `json:"idmechanic" db:"idmechanic"`
	Mechanicname    string `json:"mechanicname" db:"mechanicname"`
	Appointmentdate string `json:"appointmentdate" db:"appointmentdate"`
	Timeroutine     int64  `json:"timeroutine" db:"timeroutine"`
	Idroutine       int64  `json:"idroutine" db:"idroutine"`
}

type AppointmentsClientResponse struct {
	Idappointment   int64  `json:"idappointment" db:"idappointment"`
	Idmechanic      int64  `json:"idmechanic" db:"idmechanic"`
	Name            string `json:"name" db:"name"`
	LastName        string `json:"last_name" db:"last_name"`
	Appointmentdate string `json:"appointmentdate" db:"appointmentdate"`
	Timeroutine     string `json:"timeroutine" db:"timeroutine"`
	Idroutine       int64  `json:"idroutine" db:"idroutine"`
	Status          int64  `json:"status" db:"status"`
	Nameworkshop    string `json:"nameworkshop" db:"nameworkshop"`
}

type AppointmentsNotAvailableResponse struct {
	Idmechanic      int64  `json:"idmechanic" db:"idmechanic"`
	Name            string `json:"name" db:"name"`
	LastName        string `json:"last_name" db:"last_name"`
	Appointmentdate string `json:"appointmentdate" db:"appointmentdate"`
	Timeroutine     string `json:"timeroutine" db:"timeroutine"`
	Idroutine       string `json:"idroutine" db:"idroutine"`
}

// GetAppointmentsFilter traer todas las citas del taller filtrado por id
func GetAppointmentsFilter(id1 string, id2 string, date1 string, date2 string) ([]AppointmentsResponse, error) {
	appointmentsResponse := []AppointmentsResponse{}

	query := "select * from getappointments(" + id1 + "," + id2 + ",'" + date1 + "','" + date2 + "')"

	error := shared.GetDb().Raw(query).Find(&appointmentsResponse).Error
	return appointmentsResponse, error
}

// GetAppointmentsByClient traer todas las citas del usuario final
func GetAppointmentsByClient(userID string, fhinitial string, fhend string) ([]AppointmentsClientResponse, error) {
	appointmentsClientResponse := []AppointmentsClientResponse{}

	query := "select * from getappointmentsuser(" + userID + ",'" + fhinitial + "','" + fhend + "')"

	error := shared.GetDb().Raw(query).Find(&appointmentsClientResponse).Error
	return appointmentsClientResponse, error
}

// GetAppointmentsNotAvailables traer todas las citas no disponibles
func GetAppointmentsNotAvailables(workshopID string, routineID string, fhinitial string, fhend string) ([]AppointmentsNotAvailableResponse, error) {
	appointmentsNotAvailableResponse := []AppointmentsNotAvailableResponse{}

	query := "select * from getappointmentsnotavailables(" + workshopID + "," + routineID + ",'" + fhinitial + "','" + fhend + "')"

	error := shared.GetDb().Raw(query).Find(&appointmentsNotAvailableResponse).Error
	return appointmentsNotAvailableResponse, error
}
