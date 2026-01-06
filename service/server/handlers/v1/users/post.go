package users

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"goliath/repositories"
	"goliath/types/api"
	"goliath/types/postgres"
)

type UsersPost struct{}

func (_ UsersPost) GetPath() string {
	return "/api/v1/users"
}

func (_ UsersPost) GetMethod() string {
	return http.MethodPost
}

func (_ UsersPost) DoHandle(c echo.Context) error {
	// Если id передан - обновляем пользователя
	var u api.User
	if err := c.Bind(&u); err != nil {
		return err
	}

	if u.Id != nil {
		// Получаем текущего пользователя для сохранения полей, которые не переданы
		existingUsers, err := repositories.GetUsers(c.Request().Context(), 1, u.Id, true)
		if err != nil {
			return err
		}

		if len(existingUsers) == 0 {
			return c.JSON(http.StatusNotFound, api.Error{
				Code:    "not_found",
				Message: "User not found",
			})
		}

		existingUser := existingUsers[0]

		// Обновляем только переданные поля
		updateUser := postgres.User{
			Id:        sql.NullInt64{Valid: true, Int64: *u.Id},
			Name:      u.Name,
			Email:     u.Email,
			Password:  existingUser.Password, // Пароль не обновляем через этот endpoint
			DeletedAt: existingUser.DeletedAt,
		}

		// Если email не передан, используем существующий
		if u.Email == "" {
			updateUser.Email = existingUser.Email
		}

		// Если name не передан, используем существующий
		if u.Name == "" {
			updateUser.Name = existingUser.Name
		}

		// Если permissions не переданы, используем существующие
		if u.Permissions == nil {
			updateUser.Permissions = existingUser.Permissions
		} else {
			updateUser.Permissions = u.Permissions
		}

		_, err = repositories.UpsertUser(
			c.Request().Context(),
			updateUser,
		)

		if err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}

	// Если id не передан - создаем нового пользователя
	// Для создания нужны name, email и password
	var req api.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, api.Error{
			Code:    "bad_request",
			Message: "Invalid request body",
		})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, api.Error{
			Code:    "bad_request",
			Message: "Name, email and password are required for user creation",
		})
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api.Error{
			Code:    "internal_error",
			Message: "Failed to hash password",
		})
	}

	// Устанавливаем права по умолчанию, если не переданы
	permissions := req.Permissions
	if permissions == nil || len(permissions) == 0 {
		permissions = []string{"read:own", "write:own"}
	}

	newUser := postgres.User{
		Id:          sql.NullInt64{Valid: false},
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Permissions: permissions,
		DeletedAt:   sql.NullTime{Valid: false},
	}

	_, err = repositories.UpsertUser(
		c.Request().Context(),
		newUser,
	)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
