package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	userNotFoundOrInvalidPassword = "User not found or invalid password"
)

type Service struct {
	Users         map[string]User   // email → user
	refreshTokens map[string]string // refresh → userID
}

func NewService() *Service {
	return &Service{
		Users:         make(map[string]User),
		refreshTokens: make(map[string]string),
	}
}

func (s *Service) Register(email, password string) (*User, error) {
	_, ok := s.Users[email]
	if ok {
		return nil, errors.New("User already exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := User{
		ID:       email,
		Email:    email,
		Password: string(hash),
		Role:     "user",
	}

	s.Users[email] = user

	return &user, nil
}

func (s *Service) Login(email, password string) (*User, error) {
	user, ok := s.Users[email]
	if !ok {
		return nil, errors.New(userNotFoundOrInvalidPassword)
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	) != nil {
		return nil, errors.New(userNotFoundOrInvalidPassword)
	}

	return &user, nil
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
