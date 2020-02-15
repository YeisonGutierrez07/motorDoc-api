package vehicles

import "github.com/motorDoc-api/v1/entities"

func RegisterNewVehicle(user *entities.User, data CreateVehicle) (entities.Vehicle, error) {
	newVehicle := entities.Vehicle{}
	newVehicle.ClientID = user.ID
	newVehicle.BrandID = data.BrandID
	newVehicle.Model = data.Model
	newVehicle.Estado = data.Estado
	newVehicle.Image = data.Image
	newVehicle.Reference = data.Reference

	vehicle, err := entities.CreateVehicle(newVehicle)
	return vehicle, err
}
