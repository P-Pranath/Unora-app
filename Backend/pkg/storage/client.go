// pkg/storage/client.go
package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/logger"
)

// Client defines the interface for storage operations
type Client interface {
	// Upload uploads a file and returns the object key
	Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error)
	// Download returns a reader for the object
	Download(ctx context.Context, objectName string) (io.ReadCloser, error)
	// Delete removes an object
	Delete(ctx context.Context, objectName string) error
	// GetPresignedDownloadURL returns a presigned URL for downloading
	GetPresignedDownloadURL(ctx context.Context, objectName string) (string, error)
	// GetPresignedUploadURL returns a presigned URL for uploading
	GetPresignedUploadURL(ctx context.Context, objectName string, contentType string) (string, error)
	// GetPresignedPreviewURL returns a presigned URL for previewing (shorter expiry)
	GetPresignedPreviewURL(ctx context.Context, objectName string) (string, error)
	// GetPublicURL returns the public accessible URL (for already public objects)
	GetPublicURL(objectName string) string
}

// MinioClient implements Client interface using MinIO
type MinioClient struct {
	client        *minio.Client
	bucket        string
	prefix        string
	publicBaseURL string
	downloadExp   time.Duration
	previewExp    time.Duration
	uploadExp     time.Duration
}

// NewClientFromConfig creates a new storage client from config
func NewClientFromConfig(cfg *config.Config) (Client, error) {
	log := logger.GetLogger("storage")

	minioClient, err := minio.New(cfg.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		Secure: cfg.Storage.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, cfg.Storage.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, cfg.Storage.Bucket, minio.MakeBucketOptions{
			Region: cfg.Storage.Region,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}
		log.Info().Str("bucket", cfg.Storage.Bucket).Msg("Created storage bucket")
	}

	return &MinioClient{
		client:        minioClient,
		bucket:        cfg.Storage.Bucket,
		prefix:        cfg.Storage.Prefix,
		publicBaseURL: cfg.Storage.PublicBaseURL,
		downloadExp:   time.Duration(cfg.Storage.URLDownloadExpiryMins) * time.Minute,
		previewExp:    time.Duration(cfg.Storage.URLPreviewExpiryMins) * time.Minute,
		uploadExp:     time.Duration(cfg.Storage.URLUploadExpiryMins) * time.Minute,
	}, nil
}

// Upload uploads a file to MinIO
func (c *MinioClient) Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error) {
	fullPath := c.getFullPath(objectName)

	_, err := c.client.PutObject(ctx, c.bucket, fullPath, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload object: %w", err)
	}

	return fullPath, nil
}

// Download returns a reader for the object
func (c *MinioClient) Download(ctx context.Context, objectName string) (io.ReadCloser, error) {
	fullPath := c.getFullPath(objectName)

	object, err := c.client.GetObject(ctx, c.bucket, fullPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to download object: %w", err)
	}

	return object, nil
}

// Delete removes an object
func (c *MinioClient) Delete(ctx context.Context, objectName string) error {
	fullPath := c.getFullPath(objectName)

	err := c.client.RemoveObject(ctx, c.bucket, fullPath, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}

// GetPresignedDownloadURL returns a presigned URL for downloading
func (c *MinioClient) GetPresignedDownloadURL(ctx context.Context, objectName string) (string, error) {
	fullPath := c.getFullPath(objectName)

	reqParams := url.Values{}
	// Set response content disposition for download
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectName))

	presignedURL, err := c.client.PresignedGetObject(ctx, c.bucket, fullPath, c.downloadExp, reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	// Replace internal endpoint with public base URL if configured
	return c.replaceWithPublicURL(presignedURL), nil
}

// GetPresignedUploadURL returns a presigned URL for uploading
func (c *MinioClient) GetPresignedUploadURL(ctx context.Context, objectName string, contentType string) (string, error) {
	fullPath := c.getFullPath(objectName)

	presignedURL, err := c.client.PresignedPutObject(ctx, c.bucket, fullPath, c.uploadExp)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}

	// Replace internal endpoint with public base URL if configured
	return c.replaceWithPublicURL(presignedURL), nil
}

// GetPresignedPreviewURL returns a presigned URL for previewing (shorter expiry)
func (c *MinioClient) GetPresignedPreviewURL(ctx context.Context, objectName string) (string, error) {
	fullPath := c.getFullPath(objectName)

	reqParams := url.Values{}
	// Set response content disposition for inline viewing
	reqParams.Set("response-content-disposition", "inline")

	presignedURL, err := c.client.PresignedGetObject(ctx, c.bucket, fullPath, c.previewExp, reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned preview URL: %w", err)
	}

	// Replace internal endpoint with public base URL if configured
	return c.replaceWithPublicURL(presignedURL), nil
}

// GetPublicURL returns the public URL for an object
func (c *MinioClient) GetPublicURL(objectName string) string {
	fullPath := c.getFullPath(objectName)
	return fmt.Sprintf("%s/%s/%s", c.publicBaseURL, c.bucket, fullPath)
}

// getFullPath returns the full object path with prefix
func (c *MinioClient) getFullPath(objectName string) string {
	if c.prefix == "" {
		return objectName
	}
	return fmt.Sprintf("%s/%s", c.prefix, objectName)
}

// replaceWithPublicURL replaces the internal MinIO endpoint with the public base URL
func (c *MinioClient) replaceWithPublicURL(presignedURL *url.URL) string {
	if c.publicBaseURL == "" {
		return presignedURL.String()
	}

	publicURL, err := url.Parse(c.publicBaseURL)
	if err != nil {
		return presignedURL.String()
	}

	// Replace host with public URL host
	presignedURL.Scheme = publicURL.Scheme
	presignedURL.Host = publicURL.Host

	return presignedURL.String()
}
