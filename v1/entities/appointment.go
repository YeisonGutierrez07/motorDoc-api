package entities

import "github.com/motorDoc-api/shared"

type AppointmentsResponse struct {
	Idmechanic      int64  `json:"idmechanic" db:"idmechanic"`
	Appointmentdate string `json:"appointmentdate" db:"appointmentdate"`
	Timeroutine     int64  `json:"timeroutine" db:"timeroutine"`
	Idroutine       int64  `json:"idroutine" db:"idroutine"`
}

// GetAppointmentsFilter traer todas las citas del taller filtrado por id
func GetAppointmentsFilter(id1 string, id2 string, date1 string, date2 string) ([]AppointmentsResponse, error) {
	appointmentsResponse := []AppointmentsResponse{}

	query := "select * from getappointments(" + id1 + "," + id2 + ",'" + date1 + "','" + date2 + "')"

	error := shared.GetDb().Raw(query).Find(&appointmentsResponse).Error
	return appointmentsResponse, error
}
