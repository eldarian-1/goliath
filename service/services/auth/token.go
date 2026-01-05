package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	AccessSecret  = []byte("ACCESS_SECRET")
	RefreshSecret = []byte("REFRESH_SECRET")
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(AccessSecret)
}

func GenerateRefreshToken(user User) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(RefreshSecret)
}
