package mechanic

// CreateMechanic Modelo para el servicio de crear mecanicos
type CreateMechanic struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Credential  string `json:"credential" binding:"required"`
	ProfilePic  string `json:"profile_pic" binding:"required"`
	MobilePhone int64  `json:"mobile_phone" binding:"required"`
}
