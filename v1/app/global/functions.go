package global

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/motorDoc-api/shared"
)

// ResponseServices Función para responder a los servicios
func ResponseServices(Data interface{}, Status string, Message string) shared.Response {
	response := shared.Response{}
	response.Data = Data
	response.Status = Status
	response.Message = Message
	return response
}

// EncryptPassword encritar contraseña
func EncryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
