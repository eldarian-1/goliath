package caches

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"goliath/models/redis"
	"goliath/models/s3"
	"io"
	"time"
)

type file struct {
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	ContentDisposition string `json:"content_disposition"`
	Bytes              []byte `json:"bytes"`
}

type Files struct{}

func (_ Files) Get(ctx context.Context, filename string) (*s3.File, error) {
	marshalledFile, err := redis.Get(ctx, filename)
	if err == nil {
		fmt.Println("File found in Redis")

		cacheFile := &file{}
		err = json.Unmarshal(marshalledFile, cacheFile)
		if err != nil {
			return nil, err
		}

		return &s3.File{
			Name:               cacheFile.Name,
			ContentType:        cacheFile.ContentType,
			ContentDisposition: cacheFile.ContentDisposition,
			Reader:             bytes.NewReader(cacheFile.Bytes),
		}, nil
	}

	s3file, err := s3.Get(ctx, filename)
	if err != nil {
		return nil, err
	}

	bytesFile, err := io.ReadAll(s3file.Reader)
	if err != nil {
		return nil, err
	}

	cacheFile := &file{
		Name:               s3file.Name,
		ContentType:        s3file.ContentType,
		ContentDisposition: s3file.ContentDisposition,
		Bytes:              bytesFile,
	}

	marshalledFile, err = json.Marshal(cacheFile)
	if err != nil {
		return nil, err
	}

	err = redis.Set(ctx, filename, marshalledFile, 5*time.Minute)
	if err != nil {
		return nil, err
	}

	fmt.Println("File found in S3")

	s3file.Reader = bytes.NewReader(bytesFile)

	return s3file, nil
}
