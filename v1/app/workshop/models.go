package workshop

// CreateWorkShop Modelo para el servicio de crear talleres
type CreateWorkShop struct {
	Name            string `json:"name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Address         string `json:"address" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Credential      string `json:"credential" binding:"required"`
	ProfilePic      string `json:"profile_pic" binding:"required"`
	MobilePhone     int64  `json:"mobile_phone" binding:"required"`
	AddressWorkshop string `json:"address_workshop" binding:"required"`
	ImageWorkshop   string `json:"imageWorkshop" binding:"required"`
	NameWorkshop    string `json:"name_workshop" binding:"required"`
}
