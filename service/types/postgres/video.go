package postgres

import (
	"database/sql"
	"time"
)

type Video struct {
	Id          sql.NullInt64
	Title       string
	Description sql.NullString
	FileName    string
	FileSize    int64
	ContentType string
	Duration    sql.NullInt32
	Progress    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
