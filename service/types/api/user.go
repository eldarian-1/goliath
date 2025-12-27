package api

import (
	"time"
)

type User struct {
	Id        *int64     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GetUsersResponse struct {
	Users      []User `json:"users"`
	CursorById *int64 `json:"cursor"`
}
