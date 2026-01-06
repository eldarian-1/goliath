package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"

	"goliath/logics/auth"
	"goliath/types/api"
)

// RequirePermission creates a middleware that checks if the user has the required permission
func RequirePermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get user from context (set by JWT middleware)
			userInterface := c.Get("user")
			if userInterface == nil {
				return api.NewUnauthorized(c)
			}

			user, ok := userInterface.(*auth.User)
			if !ok {
				return api.NewUnauthorized(c)
			}

			// Check if user has the required permission
			if !hasPermission(user.Permissions, permission) {
				return api.NewForbidden(c)
			}

			return next(c)
		}
	}
}

// RequireAnyPermission creates a middleware that checks if the user has any of the required permissions
func RequireAnyPermission(permissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userInterface := c.Get("user")
			if userInterface == nil {
				return api.NewUnauthorized(c)
			}

			user, ok := userInterface.(*auth.User)
			if !ok {
				return api.NewUnauthorized(c)
			}

			// Check if user has any of the required permissions
			for _, perm := range permissions {
				if hasPermission(user.Permissions, perm) {
					return next(c)
				}
			}

			return api.NewForbidden(c)
		}
	}
}

// RequireAllPermissions creates a middleware that checks if the user has all of the required permissions
func RequireAllPermissions(permissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userInterface := c.Get("user")
			if userInterface == nil {
				return api.NewUnauthorized(c)
			}

			user, ok := userInterface.(*auth.User)
			if !ok {
				return api.NewUnauthorized(c)
			}

			// Check if user has all required permissions
			for _, perm := range permissions {
				if !hasPermission(user.Permissions, perm) {
					return api.NewForbidden(c)
				}
			}

			return next(c)
		}
	}
}

// hasPermission checks if a permission exists in the user's permissions list
// Supports wildcard permissions (e.g., "videos:*" matches "videos:read", "videos:write", etc.)
func hasPermission(userPermissions []string, required string) bool {
	for _, perm := range userPermissions {
		// Exact match
		if perm == required {
			return true
		}

		// Wildcard match (e.g., "videos:*" matches "videos:read")
		if strings.HasSuffix(perm, ":*") {
			prefix := strings.TrimSuffix(perm, "*")
			if strings.HasPrefix(required, prefix) {
				return true
			}
		}

		// Admin wildcard
		if perm == "*" || perm == "admin" {
			return true
		}
	}

	return false
}
