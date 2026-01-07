package videos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"goliath/repositories"
	"goliath/types/api"
)

type List struct{}

func (_ List) GetPath() string {
	return "/api/v1/videos"
}

func (_ List) GetMethod() string {
	return http.MethodGet
}

func (_ List) DoHandle(c echo.Context) error {
	// Get query parameters
	limitStr := c.QueryParam("limit")
	cursorStr := c.QueryParam("cursor")

	// Parse limit (default 20, max 100)
	limit := int64(20)
	if limitStr != "" {
		if parsedLimit, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			if parsedLimit > 0 && parsedLimit <= 100 {
				limit = parsedLimit
			}
		}
	}

	// Parse cursor
	var cursor *int64
	if cursorStr != "" {
		if parsedCursor, err := strconv.ParseInt(cursorStr, 10, 64); err == nil {
			cursor = &parsedCursor
		}
	}

	// Get videos from database
	videos, err := repositories.GetVideos(c.Request().Context(), limit, cursor)
	if err != nil {
		return api.NewBadRequest(c, "Failed to fetch videos")
	}

	// Convert to API response format
	response := make([]map[string]interface{}, 0, len(videos))
	for _, video := range videos {
		item := map[string]interface{}{
			"id":          video.Id.Int64,
			"title":       video.Title,
			"fileName":    video.FileName,
			"fileSize":    video.FileSize,
			"contentType": video.ContentType,
			"progress":    video.Progress,
			"createdAt":   video.CreatedAt,
		}

		if video.Description.Valid {
			item["description"] = video.Description.String
		}

		if video.Duration.Valid {
			item["duration"] = video.Duration.Int32
		}

		response = append(response, item)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"videos": response,
		"count":  len(response),
	})
}
