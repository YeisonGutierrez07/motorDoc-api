package vehicles

import "github.com/motorDoc-api/v1/entities"

func RegisterNewVehicle(user *entities.User, data CreateVehicle) (entities.Vehicle, error) {
	newVehicle := entities.Vehicle{}
	newVehicle.ClientID = user.ID
	newVehicle.Idreferencebrand = data.ReferenceID
	newVehicle.Model = data.Model
	newVehicle.Estado = data.Estado
	newVehicle.Image = data.Image
	newVehicle.Placa = data.Placa

	vehicle, err := entities.CreateVehicle(newVehicle)
	return vehicle, err
}
