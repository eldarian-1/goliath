package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
	"goliath/types/api"
)

type JWT struct{}

var skippedPaths = map[string]bool{
	"/api/v1/auth/register": true,
	"/api/v1/auth/login":    true,
}

var skippedPathPrefixes = []string{
	"/api/v1/videos/",
}

func (_ JWT) GetMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper:      skip,
		ErrorHandler: handleError,
		SigningKey:   auth.AccessSecret,
		TokenLookup:  "cookie:access",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &auth.Claims{}
		},
	})
}

func skip(c echo.Context) bool {
	path := c.Path()

	// Check exact paths
	if _, ok := skippedPaths[path]; ok {
		return true
	}

	// Check path prefixes
	for _, prefix := range skippedPathPrefixes {
		if len(path) >= len(prefix) && path[:len(prefix)] == prefix {
			return true
		}
	}

	return false
}

func handleError(c echo.Context, err error) error {
	return api.NewUnauthorized(c)
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
