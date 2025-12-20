package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"goliath/migrations"
)

func exec(query string, args ...any) (sql.Result, error) {
	db, err := sql.Open(migrations.DriverName, migrations.ConnectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	return db.Exec(query, args)
}

func query(query string, args ...any) (*sql.Rows, error) {
	db, err := sql.Open(migrations.DriverName, migrations.ConnectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	return db.Query(query, args)
}
