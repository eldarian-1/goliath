package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	AccessSecret  = []byte("ACCESS_SECRET")
	RefreshSecret = []byte("REFRESH_SECRET")
)

type Claims struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user User) (string, error) {
	claims := Claims{
		UserID:      strconv.FormatInt(user.ID, 10),
		UserName:    user.Name,
		Permissions: user.Permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(AccessSecret)
}

func GenerateRefreshToken(user User) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(user.ID, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(RefreshSecret)
}
