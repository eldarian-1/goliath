package videos

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"goliath/models/s3"
	"goliath/repositories"
	"goliath/types/api"
)

type Get struct{}

func (_ Get) GetPath() string {
	return "/api/v1/videos/:id"
}

func (_ Get) GetMethod() string {
	return http.MethodGet
}

func (_ Get) DoHandle(c echo.Context) error {
	// Get video ID from path parameter
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return api.NewBadRequest(c, "Invalid video ID")
	}

	// Get video metadata from database
	video, err := repositories.GetVideoById(c.Request().Context(), id)
	if err != nil {
		return api.NewBadRequest(c, "Video not found")
	}

	// Get video file from S3
	file, err := s3.Get(c.Request().Context(), video.FileName)
	if err != nil {
		return api.NewBadRequest(c, "Failed to retrieve video file")
	}
	defer file.Reader.(io.ReadCloser).Close()

	// Get file size
	fileSize := video.FileSize

	// Set common headers
	c.Response().Header().Set("Content-Type", video.ContentType)
	c.Response().Header().Set("Accept-Ranges", "bytes")
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")

	// Check for Range header
	rangeHeader := c.Request().Header.Get("Range")

	if rangeHeader == "" {
		// No range request - send entire file
		c.Response().Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
		c.Response().Header().Set("Content-Disposition", file.ContentDisposition)
		return c.Stream(http.StatusOK, video.ContentType, file.Reader)
	}

	// Parse Range header
	ranges := strings.TrimPrefix(rangeHeader, "bytes=")
	parts := strings.Split(ranges, "-")

	if len(parts) != 2 {
		return api.NewBadRequest(c, "Invalid Range header")
	}

	var start, end int64

	// Parse start
	if parts[0] != "" {
		start, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return api.NewBadRequest(c, "Invalid Range start")
		}
	}

	// Parse end
	if parts[1] != "" {
		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return api.NewBadRequest(c, "Invalid Range end")
		}
	} else {
		end = fileSize - 1
	}

	// Validate range
	if start < 0 || start >= fileSize || end >= fileSize || start > end {
		c.Response().Header().Set("Content-Range", fmt.Sprintf("bytes */%d", fileSize))
		return c.NoContent(http.StatusRequestedRangeNotSatisfiable)
	}

	// Calculate content length
	contentLength := end - start + 1

	// Skip to start position
	if start > 0 {
		_, err = io.CopyN(io.Discard, file.Reader, start)
		if err != nil {
			return api.NewBadRequest(c, "Failed to seek to start position")
		}
	}

	// Set headers for partial content
	c.Response().Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	c.Response().Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
	c.Response().Header().Set("Content-Disposition", file.ContentDisposition)
	c.Response().WriteHeader(http.StatusPartialContent)

	// Stream the requested range
	_, err = io.CopyN(c.Response().Writer, file.Reader, contentLength)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}
