package middlewares

import (
	"goliath/server/handlers/v1/auth/gpt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWT struct{}

var skippedPaths = map[string]bool{
	"/api/v1/auth/register": true,
	"/api/v1/auth/login":    true,
}

func (_ JWT) GetMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper:     skip,
		SigningKey:  gpt.AccessSecret,
		TokenLookup: "cookie:access",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &gpt.Claims{}
		},
	})
}

func skip(c echo.Context) bool {
	_, ok := skippedPaths[c.Path()]
	return ok
}

// func RequireRole(role string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			user := c.Get("user").(*jwt.Token)
// 			claims := user.Claims.(*gpt.Claims)

// 			if claims.Role != role {
// 				return echo.ErrForbidden
// 			}
// 			return next(c)
// 		}
// 	}
// }
