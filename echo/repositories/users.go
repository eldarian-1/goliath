package repositories

import (
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
			users
		WHERE
			($2 IS NULL OR id <= $2) AND
			($3 OR deleted_at IS NULL)
		ORDER BY id DESC
		LIMIT $1;
	`

	insertQuery = `
		INSERT INTO	users (name)
		VALUES ($1);
	`

	updateQuery = `
		UPDATE
			users
		SET
			name = $2,
			deleted_at = $3,
			updated_at = NOW()
		WHERE
			id = $1 AND
			(
				name IS DISTINCT FROM $2 OR
				(deleted_at IS NULL) IS DISTINCT FROM ($3 IS NULL)
			);
	`
)

func GetUsers(limit int64, cursorById *int64, withDeleted bool) ([]postgres.User, error) {
	rows, err := query(getQuery, cursorById, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []postgres.User{}

    for rows.Next(){
        u := postgres.User{}
        err := rows.Scan(&u.Id, &u.Name, &u.CreatedAt, &u.UpdatedAt, u.DeletedAt)
        if err != nil{
            fmt.Println(err)
            continue
        }
        users = append(users, u)
    }

	return users, nil
}

func UpsertUser(user postgres.User) error {
	if user.Id == nil {
		_, err := exec(insertQuery, user.Name)

		return err
	}

	_, err := exec(updateQuery, user.Id, user.Name, user.DeletedAt)

	return err
}
