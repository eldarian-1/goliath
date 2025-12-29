package s3

import (
	"context"
	"fmt"
	"goliath/utils"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
    bucket = "goliath"
)

var s3StorageUrl string
var s3StorageRegion string
var s3StorageKey string
var s3StorageSecret string

type File struct {
    Name        string
    ContentType string
    ContentDisposition string
    Reader      io.Reader
}

func init() {
	s3StorageUrl = utils.GetEnv("S3_STORAGE_URL", "http://localhost:9000")
	s3StorageRegion = utils.GetEnv("S3_STORAGE_REGION", "us-east-1")
	s3StorageKey = utils.GetEnv("S3_STORAGE_KEY", "minio")
	s3StorageSecret = utils.GetEnv("S3_STORAGE_SECRET", "minio123")

    client, err := getClient(context.Background())
    if err != nil {
        panic(err)
    }

    _, err = client.CreateBucket(context.Background(), &s3.CreateBucketInput{Bucket: aws.String(bucket)})
    if err != nil {
        fmt.Printf("S3 Bucket creation failed: %s", err.Error())
        return
    }

    fmt.Println("S3 Bucket successfully created")
}

func Get(ctx context.Context, fileName string) (*File, error) {
    client, err := getClient(ctx)
    if err != nil {
        return nil, err
    }

    res, err := client.GetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(fileName),
    })
    if err != nil {
        return nil, err
    }

    return &File{
        Name: fileName,
        ContentType: *res.ContentType,
        ContentDisposition: *res.ContentDisposition,
        Reader: res.Body,
    }, nil
}

func Put(ctx context.Context, file *File) error {
    client, err := getClient(ctx)
    if err != nil {
        return fmt.Errorf("Failed to get client: %s", err)
    }

    tmp, _ := os.CreateTemp("", "upload-*")
	io.Copy(tmp, file.Reader)
    tmp.Seek(0, 0)

    _, err = client.PutObject(ctx, &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(file.Name),
        Body:   tmp,
        ContentType:       aws.String(file.ContentType),
        ContentDisposition: aws.String(file.ContentDisposition),
    })

    if err != nil {
        return fmt.Errorf("Failed to put object: %s", err)
    }

    return err
}

func Delete(ctx context.Context, fileName string) error {
    client, err := getClient(ctx)
    if err != nil {
        return err
    }

    _, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(fileName),
    })

    return err
}

func getClient(ctx context.Context) (*s3.Client, error) {
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(s3StorageRegion),
        config.WithCredentialsProvider(
            credentials.NewStaticCredentialsProvider(
                s3StorageKey,
                s3StorageSecret,
                "",
            ),
        ),
        config.WithEndpointResolver(
            aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
                return aws.Endpoint{
                    URL:               s3StorageUrl,
                    SigningRegion:     s3StorageRegion,
                    HostnameImmutable: true,
                }, nil
            }),
        ),
    )

    if err != nil {
        return nil, err
    }

    return s3.NewFromConfig(cfg), nil
}
