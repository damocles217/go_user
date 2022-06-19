package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func JwtParse(mapping map[string]any) (string, error) {
	// JWT usage
	// Declare secret
	secret := os.Getenv("SECRET_JWT")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": mapping["_id"].(string),
	})

	// String token
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
