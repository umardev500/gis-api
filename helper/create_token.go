package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(claims jwt.MapClaims) (*string, error) {
	secret := os.Getenv("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generated encoded token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
