package entities

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SigningKey valirable
var SigningKey = "$SolidSigningKey$"

// BaseModel modelo base para la creación de las tablas
type BaseModel struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// EncryptPassword encritar contraseña
func EncryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}

// GenerateToken generar token en el login
func GenerateToken(key []byte, userID int64, credential string) (string, error) {

	//new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["user_id"] = userID
	claims["credential"] = credential
	claims["exp"] = time.Now().Add(time.Hour*720).UnixNano() / int64(time.Millisecond)
	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString(key)
	return tokenString, err
}
