package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	connectionUrl = "postgres://user:password@localhost:5432/goliath"
)

func Exec(query string, args ...any) (pgconn.CommandTag, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connectionUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	return conn.Exec(ctx, query, args...)
}

func Query(query string, args ...any) (pgx.Rows, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connectionUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	return conn.Query(ctx, query, args...)
}
