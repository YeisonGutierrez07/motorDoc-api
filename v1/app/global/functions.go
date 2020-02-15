package global

import (
	"github.com/YeisonGutierrez07/motorDoc-api/shared"
)

// ResponseServices Función para responder a los servicios
func ResponseServices(Data interface{}, Status string, Message string) shared.Response {
	response := shared.Response{}
	response.Data = Data
	response.Status = Status
	response.Message = Message
	return response
}
