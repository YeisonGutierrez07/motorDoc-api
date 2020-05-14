package vehicles

// Vehicle Modelo para los vehiculos
type CreateVehicle struct {
	ReferenceID int    `json:"idreferencebrand" binding:"required"`
	Model       string `json:"model" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Estado      string `json:"estado" binding:"required"`
	Placa       string `json:"placa" binding:"required"`
}
