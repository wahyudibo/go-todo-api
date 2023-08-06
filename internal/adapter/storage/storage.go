package storage

import (
	"context"
	"io"
)

//go:generate mockery --name StorageAdapter

// StorageAdapter defines methods for storage adapter.
type StorageAdapter interface {
	// Upload uploads file to the storage
	Upload(ctx context.Context, directory, objectKey string, body io.Reader) (string, error)
	// GenerateDownloadURL generates downloadable URL of the object
	GenerateDownloadURL(ctx context.Context, objectKey string) (string, error)
}
