package repositories

import (
	"context"
	"fmt"
	"goliath/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var connectionUrl string

func init() {
	connectionUrl = fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		utils.GetEnv("POSTGRES_USER", "user"),
		utils.GetEnv("POSTGRES_PASSWORD", "password"),
		utils.GetEnv("POSTGRES_HOST", "localhost:5432"),
		utils.GetEnv("POSTGRES_DB", "goliath"),
	)
}

func Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	withTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(withTimeout, connectionUrl)
	if err != nil {
		cancel()
		return pgconn.CommandTag{}, err
	}
	defer conn.Close(withTimeout)

	return conn.Exec(withTimeout, query, args...)
}

func Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	withTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(withTimeout, connectionUrl)
	if err != nil {
		cancel()
		return nil, err
	}
	defer conn.Close(withTimeout)

	return conn.Query(withTimeout, query, args...)
}
