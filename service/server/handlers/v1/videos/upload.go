package videos

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"goliath/models/s3"
	"goliath/repositories"
	"goliath/types/api"
	"goliath/types/postgres"
)

type Upload struct{}

func (_ Upload) GetPath() string {
	return "/api/v1/videos/upload"
}

func (_ Upload) GetMethod() string {
	return http.MethodPost
}

func (_ Upload) DoHandle(c echo.Context) error {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return api.NewBadRequest(c, "Failed to parse multipart form")
	}

	// Get video file
	files := form.File["video"]
	if len(files) == 0 {
		return api.NewBadRequest(c, "No video file provided")
	}

	fileHeader := files[0]

	// Validate file type
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "video/") {
		return api.NewBadRequest(c, "File must be a video")
	}

	// Open uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return api.NewBadRequest(c, "Failed to open uploaded file")
	}
	defer src.Close()

	// Get metadata
	title := c.FormValue("title")
	if title == "" {
		return api.NewBadRequest(c, "Title is required")
	}
	description := c.FormValue("description")

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("videos/%s%s", generateUniqueID(), ext)

	// Create temp file for S3 upload
	tmpFile, err := os.CreateTemp("", "video-upload-*"+ext)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy uploaded file to temp file
	_, err = io.Copy(tmpFile, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// Reset file pointer
	tmpFile.Seek(0, 0)

	// Upload to S3
	s3File := &s3.File{
		Name:               fileName,
		ContentType:        contentType,
		ContentDisposition: fmt.Sprintf("inline; filename=\"%s\"", fileHeader.Filename),
		Reader:             tmpFile,
	}

	err = s3.Put(c.Request().Context(), s3File)
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Save video metadata to database
	video := postgres.Video{
		Title:       title,
		Description: sql.NullString{String: description, Valid: description != ""},
		FileName:    fileName,
		FileSize:    fileHeader.Size,
		ContentType: contentType,
	}

	videoId, err := repositories.InsertVideo(c.Request().Context(), video)
	if err != nil {
		return fmt.Errorf("failed to save video metadata: %w", err)
	}

	// Return success response with metadata
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":     true,
		"id":          videoId,
		"fileName":    fileName,
		"title":       title,
		"description": description,
		"size":        fileHeader.Size,
		"contentType": contentType,
	})
}

func generateUniqueID() string {
	// Simple unique ID generation using timestamp
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), os.Getpid())
}
