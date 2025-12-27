package repositories

import (
	"context"
	"fmt"

	"goliath/types/postgres"
)

const (
	getQuery = `
		SELECT
			id,
			name,
			created_at,
			updated_at,
			deleted_at
		FROM
			goliath.users
		WHERE
			($2::BIGINT IS NULL OR id <= $2::BIGINT) AND
			($3::BOOLEAN OR deleted_at IS NULL)
		ORDER BY id DESC
		LIMIT $1::BIGINT;
	`

	insertQuery = `
		INSERT INTO	goliath.users (name)
		VALUES ($1::TEXT);
	`

	updateQuery = `
		UPDATE
			goliath.users
		SET
			name = $2::TEXT,
			deleted_at = $3::TIMESTAMPTZ,
			updated_at = NOW()
		WHERE
			id = $1::BIGINT AND
			(
				name IS DISTINCT FROM $2::TEXT OR
				(deleted_at IS NULL) IS DISTINCT FROM ($3::TIMESTAMPTZ IS NULL)
			);
	`
)

func GetUsers(ctx context.Context, limit int64, cursorById *int64, withDeleted bool) ([]postgres.User, error) {
	rows, err := Query(
		ctx,
		getQuery,
		limit,
		cursorById,
		withDeleted,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []postgres.User{}

	for rows.Next() {
		u := postgres.User{}
		err := rows.Scan(&u.Id, &u.Name, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func UpsertUser(ctx context.Context, user postgres.User) (bool, error) {
	if !user.Id.Valid {
		tag, err := Exec(ctx, insertQuery, user.Name)

		return tag.RowsAffected() > 0, err
	}

	tag, err := Exec(ctx, updateQuery, user.Id, user.Name, user.DeletedAt)

	return tag.RowsAffected() > 0, err
}
