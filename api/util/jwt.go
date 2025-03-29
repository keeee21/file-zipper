package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	UserID    int    `json:"user_id"`
	GoogleSub string `json:"google_sub"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, googleSub string) (string, error) {
	claims := &Claims{
		UserID:    userID,
		GoogleSub: googleSub,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "file-zipper-api",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
