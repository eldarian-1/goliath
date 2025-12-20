package postgres

import (
	"time"
)

type User struct {
	Id			*int64
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
}
