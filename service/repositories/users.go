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
			email,
			password,
			permissions,
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

	getUserByEmailQuery = `
		SELECT
			id,
			name,
			email,
			password,
			permissions,
			created_at,
			updated_at,
			deleted_at
		FROM
			goliath.users
		WHERE
			email = $1::TEXT AND
			deleted_at IS NULL;
	`

	getUserByIdQuery = `
		SELECT
			id,
			name,
			email,
			password,
			permissions,
			created_at,
			updated_at,
			deleted_at
		FROM
			goliath.users
		WHERE
			id = $1::BIGINT AND
			deleted_at IS NULL;
	`

	insertQuery = `
		INSERT INTO	goliath.users (name, email, password, permissions)
		VALUES ($1::TEXT, $2::TEXT, $3::TEXT, $4::TEXT[])
		RETURNING id;
	`

	updateQuery = `
		UPDATE
			goliath.users
		SET
			name = $2::TEXT,
			email = $3::TEXT,
			permissions = $4::TEXT[],
			deleted_at = $5::TIMESTAMPTZ,
			updated_at = NOW()
		WHERE
			id = $1::BIGINT AND
			(
				name IS DISTINCT FROM $2::TEXT OR
				email IS DISTINCT FROM $3::TEXT OR
				permissions IS DISTINCT FROM $4::TEXT[] OR
				(deleted_at IS NULL) IS DISTINCT FROM ($5::TIMESTAMPTZ IS NULL)
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
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Permissions, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUserByEmail(ctx context.Context, email string) (*postgres.User, error) {
	row := QueryRow(ctx, getUserByEmailQuery, email)

	u := postgres.User{}
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Permissions, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUserById(ctx context.Context, id int64) (*postgres.User, error) {
	row := QueryRow(ctx, getUserByIdQuery, id)

	u := postgres.User{}
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Permissions, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UpsertUser(ctx context.Context, user postgres.User) (bool, error) {
	if !user.Id.Valid {
		var id int64
		err := QueryRow(ctx, insertQuery, user.Name, user.Email, user.Password, user.Permissions).Scan(&id)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	tag, err := Exec(ctx, updateQuery, user.Id, user.Name, user.Email, user.Permissions, user.DeletedAt)

	return tag.RowsAffected() > 0, err
}
