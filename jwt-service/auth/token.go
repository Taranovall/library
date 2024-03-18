package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	jwtDuration = 180 * time.Minute
	secretKey   = "46c479974a7050f2b9647ec3f699faee43fa2de1fe58aaf6aae6c35b5ef86c9b"
)

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: username,
	})

	return token.SignedString([]byte(secretKey))
}
