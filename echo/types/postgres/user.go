package postgres

import (
	"database/sql"
	"time"
)

type User struct {
	Id        sql.NullInt64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
