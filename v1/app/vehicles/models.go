package vehicles

// Vehicle Modelo para los vehiculos
type CreateVehicle struct {
	BrandID   int64  `json:"brand_id" binding:"required"`
	Model     string `json:"model" binding:"required"`
	Image     string `json:"image" binding:"required"`
	Estado    string `json:"estado" binding:"required"`
	Reference string `json:"reference" binding:"required"`
	Placa     string `json:"placa" binding:"required"`
}
