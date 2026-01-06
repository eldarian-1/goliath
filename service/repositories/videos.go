package repositories

import (
	"context"
	"fmt"

	"goliath/types/postgres"
)

const (
	getVideosQuery = `
		SELECT
			id,
			title,
			description,
			file_name,
			file_size,
			content_type,
			duration,
			created_at,
			updated_at,
			deleted_at
		FROM
			goliath.videos
		WHERE
			($2::BIGINT IS NULL OR id <= $2::BIGINT) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1::BIGINT;
	`

	getVideoByIdQuery = `
		SELECT
			id,
			title,
			description,
			file_name,
			file_size,
			content_type,
			duration,
			created_at,
			updated_at,
			deleted_at
		FROM
			goliath.videos
		WHERE
			id = $1::BIGINT AND
			deleted_at IS NULL;
	`

	insertVideoQuery = `
		INSERT INTO goliath.videos (title, description, file_name, file_size, content_type, duration)
		VALUES ($1::TEXT, $2::TEXT, $3::TEXT, $4::BIGINT, $5::TEXT, $6::INTEGER)
		RETURNING id;
	`

	updateVideoQuery = `
		UPDATE
			goliath.videos
		SET
			title = $2::TEXT,
			description = $3::TEXT,
			duration = $4::INTEGER,
			updated_at = NOW()
		WHERE
			id = $1::BIGINT AND
			deleted_at IS NULL;
	`

	deleteVideoQuery = `
		UPDATE
			goliath.videos
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE
			id = $1::BIGINT AND
			deleted_at IS NULL;
	`
)

func GetVideos(ctx context.Context, limit int64, cursorById *int64) ([]postgres.Video, error) {
	rows, err := Query(
		ctx,
		getVideosQuery,
		limit,
		cursorById,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	videos := []postgres.Video{}

	for rows.Next() {
		v := postgres.Video{}
		err := rows.Scan(
			&v.Id,
			&v.Title,
			&v.Description,
			&v.FileName,
			&v.FileSize,
			&v.ContentType,
			&v.Duration,
			&v.CreatedAt,
			&v.UpdatedAt,
			&v.DeletedAt,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		videos = append(videos, v)
	}

	return videos, nil
}

func GetVideoById(ctx context.Context, id int64) (*postgres.Video, error) {
	rows, err := Query(ctx, getVideoByIdQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("video not found")
	}

	v := postgres.Video{}
	err = rows.Scan(
		&v.Id,
		&v.Title,
		&v.Description,
		&v.FileName,
		&v.FileSize,
		&v.ContentType,
		&v.Duration,
		&v.CreatedAt,
		&v.UpdatedAt,
		&v.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func InsertVideo(ctx context.Context, video postgres.Video) (int64, error) {
	rows, err := Query(
		ctx,
		insertVideoQuery,
		video.Title,
		video.Description,
		video.FileName,
		video.FileSize,
		video.ContentType,
		video.Duration,
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("failed to insert video")
	}

	var id int64
	err = rows.Scan(&id)
	return id, err
}

func UpdateVideo(ctx context.Context, video postgres.Video) (bool, error) {
	tag, err := Exec(
		ctx,
		updateVideoQuery,
		video.Id,
		video.Title,
		video.Description,
		video.Duration,
	)

	return tag.RowsAffected() > 0, err
}

func DeleteVideo(ctx context.Context, id int64) (bool, error) {
	tag, err := Exec(ctx, deleteVideoQuery, id)

	return tag.RowsAffected() > 0, err
}
