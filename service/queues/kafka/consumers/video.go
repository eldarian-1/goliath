package consumers

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"goliath/models/s3"
	"goliath/queues/kafka/messages"
	"goliath/repositories"
)

type Video struct{}

func (_ Video) GetTopic() string {
	return "video-processing"
}

func (_ Video) Process(message []byte) error {
	var videoMsg messages.Video
	err := json.Unmarshal(message, &videoMsg)
	if err != nil {
		return errors.New("Deserializing of video message was failed")
	}

	return processVideo(context.Background(), videoMsg)
}

func processVideo(ctx context.Context, videoMsg messages.Video) error {
	fmt.Printf("Processing video %d: %s\n", videoMsg.VideoId, videoMsg.FileName)

	// Get video from database
	_, err := repositories.GetVideoById(ctx, videoMsg.VideoId)
	if err != nil {
		return fmt.Errorf("failed to get video from database: %w", err)
	}

	// Get original video file from S3
	s3File, err := s3.Get(ctx, videoMsg.FileName)
	if err != nil {
		return fmt.Errorf("failed to get video from S3: %w", err)
	}
	if closer, ok := s3File.Reader.(io.ReadCloser); ok {
		defer closer.Close()
	}

	// Create temporary directory for processing
	tmpDir, err := os.MkdirTemp("", "video-processing-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Save original file to temp directory
	inputPath := filepath.Join(tmpDir, "input"+filepath.Ext(videoMsg.FileName))
	inputFile, err := os.Create(inputPath)
	if err != nil {
		return fmt.Errorf("failed to create input file: %w", err)
	}

	_, err = io.Copy(inputFile, s3File.Reader)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// Get video duration using ffprobe
	duration, err := getDuration(inputPath)
	if err != nil {
		return fmt.Errorf("failed to get video duration: %w", err)
	}

	// Convert video to MP4 using ffmpeg with progress tracking
	outputPath := filepath.Join(tmpDir, "output.mp4")
	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		"-i", inputPath,
		"-c:v", "libx264",
		"-c:a", "aac",
		"-movflags", "+faststart",
		"-progress", "pipe:1",
		"-nostats",
		"-y", // Overwrite output file
		outputPath,
	)

	// Capture stderr for ffmpeg logs
	cmd.Stderr = os.Stderr

	// Capture stdout for progress
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	// Track progress
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "out_time_ms=") {
			msStr := strings.TrimPrefix(line, "out_time_ms=")
			ms, err := strconv.ParseFloat(msStr, 64)
			if err != nil {
				continue
			}

			// Calculate progress percentage
			seconds := ms / 1_000_000.0
			percent := (seconds / duration) * 100.0
			if percent > 100 {
				percent = 100
			}

			// Update progress in database
			progress := int(percent)
			_, err = repositories.UpdateVideoProgress(ctx, videoMsg.VideoId, progress)
			if err != nil {
				fmt.Printf("Failed to update progress: %v\n", err)
			} else {
				fmt.Printf("Video %d progress: %d%%\n", videoMsg.VideoId, progress)
			}
		}
	}

	// Wait for command to complete
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("ffmpeg conversion failed: %w", err)
	}

	// Set progress to 100% after successful conversion
	_, err = repositories.UpdateVideoProgress(ctx, videoMsg.VideoId, 100)
	if err != nil {
		fmt.Printf("Failed to set final progress: %v\n", err)
	}

	// Get converted file info
	outputFile, err := os.Open(outputPath)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer outputFile.Close()

	outputFileInfo, err := outputFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to get output file info: %w", err)
	}

	// Generate new filename for converted video
	newFileName := fmt.Sprintf("videos/%d.mp4", videoMsg.VideoId)

	// Upload converted video to S3
	outputFile.Seek(0, 0)
	s3UploadFile := &s3.File{
		Name:               newFileName,
		ContentType:        "video/mp4",
		ContentDisposition: fmt.Sprintf("inline; filename=\"%d.mp4\"", videoMsg.VideoId),
		Reader:             outputFile,
	}

	err = s3.Put(ctx, s3UploadFile)
	if err != nil {
		return fmt.Errorf("failed to upload converted video to S3: %w", err)
	}

	// Update video record in database with new file info
	_, err = repositories.UpdateVideoFile(
		ctx,
		videoMsg.VideoId,
		newFileName,
		outputFileInfo.Size(),
		"video/mp4",
		100, // progress = 100
	)
	if err != nil {
		return fmt.Errorf("failed to update video in database: %w", err)
	}

	fmt.Printf("Successfully processed video %d: %s -> %s\n", videoMsg.VideoId, videoMsg.FileName, newFileName)
	return nil
}

func getDuration(path string) (float64, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		path,
	)

	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
}
