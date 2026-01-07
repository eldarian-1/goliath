package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"goliath/repositories"
	"goliath/types/postgres"
)

const (
	userNotFoundOrInvalidPassword = "User not found or invalid password"
)

type Service struct {
	refreshTokens map[string]string // refresh → userID (string representation of int64)
}

func NewService() *Service {
	return &Service{
		refreshTokens: make(map[string]string),
	}
}

func (s *Service) GetUser(ctx context.Context, userID string) (*User, bool) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, false
	}

	pgUser, err := repositories.GetUserById(ctx, id)
	if err != nil {
		return nil, false
	}

	if !pgUser.Id.Valid {
		return nil, false
	}

	return &User{
		ID:          pgUser.Id.Int64,
		Name:        pgUser.Name,
		Email:       pgUser.Email,
		Password:    pgUser.Password,
		Permissions: pgUser.Permissions,
	}, true
}

func (s *Service) Register(ctx context.Context, name, email, password string, permissions []string) (*User, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := repositories.GetUserByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, errors.New("User already exists")
	}

	// Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Default permissions for new users
	if len(permissions) == 0 {
		permissions = []string{"read:own", "write:own"}
	}

	// Создаем пользователя в postgres
	pgUser := postgres.User{
		Id:          sql.NullInt64{Valid: false},
		Name:        name,
		Email:       email,
		Password:    string(hash),
		Permissions: permissions,
		DeletedAt:   sql.NullTime{Valid: false},
	}

	_, err = repositories.UpsertUser(ctx, pgUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Получаем созданного пользователя
	createdUser, err := repositories.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get created user: %w", err)
	}

	if !createdUser.Id.Valid {
		return nil, errors.New("failed to get user id after creation")
	}

	return &User{
		ID:          createdUser.Id.Int64,
		Name:        createdUser.Name,
		Email:       createdUser.Email,
		Password:    createdUser.Password,
		Permissions: createdUser.Permissions,
	}, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*User, error) {
	pgUser, err := repositories.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New(userNotFoundOrInvalidPassword)
	}

	if !pgUser.Id.Valid {
		return nil, errors.New(userNotFoundOrInvalidPassword)
	}

	// Проверяем пароль
	if bcrypt.CompareHashAndPassword(
		[]byte(pgUser.Password),
		[]byte(password),
	) != nil {
		return nil, errors.New(userNotFoundOrInvalidPassword)
	}

	return &User{
		ID:          pgUser.Id.Int64,
		Name:        pgUser.Name,
		Email:       pgUser.Email,
		Password:    pgUser.Password,
		Permissions: pgUser.Permissions,
	}, nil
}

func (s *Service) SaveRefresh(token, userID string) {
	s.refreshTokens[token] = userID
}

func (s *Service) DeleteRefresh(token string) {
	delete(s.refreshTokens, token)
}

func (s *Service) ValidateRefresh(token string) (string, bool) {
	id, ok := s.refreshTokens[token]
	return id, ok
}
