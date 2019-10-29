package middlewares

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/shared"
	"github.com/motorDoc-api/v1/app/users"
)

// SigningKey valirable
var SigningKey = "$SolidSigningKey$"

// AuthHandler agregar las rutas a un validador
func AuthHandler(authRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Su solicitud no esta autorizada", "status": shared.StatusError, "data": nil})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "No se proporcionó un token de autorización", "status": shared.StatusError, "data": nil})
			c.Abort()
			return
		}

		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(403, gin.H{"message": "Token de autorización inválido", "status": shared.StatusError, "data": nil})
			c.Abort()
			return
		}

		user := &users.User{}
		if !shared.GetDb().Set("gorm:auto_preload", true).Where("id = ?", valid.Claims.(jwt.MapClaims)["user_id"]).First(&user).RecordNotFound() {
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(403, gin.H{"message": "Token de autorización inválido", "status": shared.StatusError, "data": nil})
			c.Abort()
			return
		}
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
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

// ValidateToken Validar token enviado desde el cliente
func ValidateToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}
