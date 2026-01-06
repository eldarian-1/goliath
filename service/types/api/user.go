package api

import (
	"time"
)

type User struct {
	Id          *int64     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Permissions []string   `json:"permissions"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type GetUsersResponse struct {
	Users      []User `json:"users"`
	CursorById *int64 `json:"cursor"`
}

type CreateUserRequest struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Permissions []string `json:"permissions,omitempty"`
}

type UpdateUserRequest struct {
	Name        *string  `json:"name,omitempty"`
	Email       *string  `json:"email,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}
